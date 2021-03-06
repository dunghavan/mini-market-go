package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
	"time"
)

type User struct {
	Id          int64        `orm:"column(id);auto" json:"id"`
	Login       string       `orm:"column(login);size(128)" json:"login"`
	Password    string       `orm:"column(password);size(128);null" json:"password"`
	Name        string       `orm:"column(name);size(128);null" json:"name"`
	FirstName   string       `orm:"column(first_name);size(128);null" json:"firstName"`
	LastName    string       `orm:"column(last_name);size(128);null" json:"lastName"`
	Email       string       `orm:"column(email);size(128);null" json:"email"`
	Phone       string       `orm:"column(phone);size(128);null" json:"phone"`
	ImageUrl    string       `orm:"column(image_url);size(128);null" json:"imageUrl"`
	Activated   bool         `orm:"column(activated);null" json:"activated"`
	CreatedDate time.Time    `orm:"column(created_date);size(128);null" json:"createdDate"`
	FacebookId  string       `orm:"column(facebook_id);size(128);null" json:"facebookId"`
	Authorities []*Authority `orm:"rel(m2m);rel_table(user_authority)" json:"authorities"`
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func CreateOrUpdate(fb *FaceBookUser) (res *User, err error) {
	m := &User{
		Name:       fb.Name,
		FirstName:  fb.FirstName,
		LastName:   fb.LastName,
		Email:      fb.Email,
		ImageUrl:   fb.Picture.Data.Url,
		FacebookId: fb.Id,
	}
	o := orm.NewOrm()
	var u User
	if err = o.QueryTable(new(User)).Filter("email", m.Email).Filter("facebook_id", m.FacebookId).One(&u); err != nil {
		// Insert
		if id, err := o.Insert(m); err != nil {
			glog.Errorf("Insert new User with facebook_id=%s error: %s", m.FacebookId, err.Error())
			return nil, err
		} else {
			glog.Infof("Insert new User with facebook_id=%s success", m.FacebookId)
			m.Id = id
			return m, nil
		}
	} else {
		// Update existing User
		if userNotChange(m, &u) {
			glog.Info("User info not change...")
		} else {
			m.Id = u.Id
			if n, err := o.Update(m); err != nil {
				glog.Errorf("Update User with facebook_id=%s error: ", m.FacebookId, err.Error())
				return nil, err
			} else {
				glog.Infof("Update %v row(s) User with facebook_id=%s success", n, m.FacebookId)
			}
		}
		return m, nil
	}
}

func userNotChange(a, b *User) bool {
	return a.Name == b.Name && strings.Split(a.ImageUrl, "&")[0] == strings.Split(b.ImageUrl, "&")[0]
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		if _, err = o.LoadRelated(v, "Authorities"); err != nil && err != orm.ErrNoRows {
			glog.Errorf("Query User LoadRelated Authorities error: %s", err.Error())
		}
		return v, nil
	}
	return nil, err
}

func GetUserByUsernamePassword(username, password string) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Login: username, Password: password}
	if err = o.QueryTable(new(User)).Filter("Login", username).Filter("Password", password).RelatedSel().One(v); err == nil {
		if _, err = o.LoadRelated(v, "Authorities"); err != nil && err != orm.ErrNoRows {
			glog.Errorf("Query User LoadRelated Authorities error: %s", err.Error())
		}
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []User
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
