package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
)

type Authority struct {
	Id    int64   `orm:"column(id);auto" json:"id"`
	Name  string  `orm:"column(name);size(128)" json:"name"`
	Users []*User `orm:"reverse(many)" json:"users"`
}

func init() {
	orm.RegisterModel(new(Authority))
}

// AddAuthority insert a new Authority into database and returns
// last inserted Id on success.
func AddAuthority(m *Authority) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	if err == nil {
		m.Id = id
		m2m := o.QueryM2M(m, "Users")
		if _, err = m2m.Add(m.Users[0]); err != nil {
			glog.Errorf("Add m2m users err: %s", err.Error())
		}
	}
	return
}

// GetAuthorityById retrieves Authority by Id. Returns error if
// Id doesn't exist
func GetAuthorityById(id int64) (v *Authority, err error) {
	o := orm.NewOrm()
	v = &Authority{Id: id}
	if err = o.QueryTable(new(Authority)).Filter("Id", id).RelatedSel().One(v); err == nil {
		if _, err = o.LoadRelated(v, "Users"); err != nil && err != orm.ErrNoRows {
			glog.Errorf("Load Authority Related Users error: %s", err.Error())
		}
		return v, nil
	}
	return nil, err
}

// GetAllAuthority retrieves all Authority matches certain condition. Returns empty list if
// no records exist
func GetAllAuthority(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Authority))
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

	var l []Authority
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

// UpdateAuthority updates Authority by Id and returns error if
// the record to be updated doesn't exist
func UpdateAuthorityById(m *Authority) (err error) {
	o := orm.NewOrm()
	v := Authority{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAuthority deletes Authority by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAuthority(id int64) (err error) {
	o := orm.NewOrm()
	v := Authority{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Authority{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
