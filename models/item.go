package models

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Item struct {
	Id          int64     `orm:"column(id);auto" json:"id"`
	Name        string    `orm:"column(name);size(255);null" json:"name"`
	Desc        string    `orm:"column(desc);size(255);null" json:"desc"`
	Price       float64   `orm:"column(price);default(0)" json:"price"`
	IsAvailable bool      `orm:"column(is_available);default(true)" json:"isAvailable"`
	DeliveryWay string    `orm:"column(delivery_way);size(255);null" json:"deliveryWay"`
	Address     string    `orm:"column(address);size(255);null" json:"address"`
	Note        string    `orm:"column(note);size(255);null" json:"note"`
	State       bool      `orm:"column(state);default(true)" json:"state"`
	Phone       string    `orm:"column(phone);size(128);null" json:"phone"`
	CreatedDate time.Time `orm:"column(created_date);type(time_stamp);auto_now;null" json:"createdDate"`
	Type        *Type     `orm:"column(type);rel(fk);null;on_delete(set_null)" json:"type"`
	User        *User     `orm:"column(user);rel(fk);null;on_delete(set_null)" json:"user"`
	Images      []*Image  `orm:"-" json:"images"` // For response to client
}

func init() {
	orm.RegisterModel(new(Item))
}

// AddItem insert a new Item into database and returns
// last inserted Id on success.
func AddItem(m *Item) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetItemById retrieves Item by Id. Returns error if
// Id doesn't exist
func GetItemById(id int64) (v *Item, err error) {
	o := orm.NewOrm()
	v = &Item{Id: id}
	if err = o.QueryTable(new(Item)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllItem retrieves all Item matches certain condition. Returns empty list if
// no records exist
func GetAllItem(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Item))
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
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
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
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Item
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				imgs, err := GetImageByItemId(int(v.Id))
				if err != nil {
					glog.Errorf("Get all item relate to images err: %s", err.Error())
				} else {
					v.Images = imgs
				}
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
		total, _ := qs.Count()
		return ml, total, nil
	}
	return nil, 0, err
}

func GetImage() {

}

// UpdateItem updates Item by Id and returns error if
// the record to be updated doesn't exist
func UpdateItemById(m *Item) (err error) {
	o := orm.NewOrm()
	v := Item{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteItem deletes Item by Id and returns error if
// the record to be deleted doesn't exist
func DeleteItem(id int64) (err error) {
	o := orm.NewOrm()
	v := Item{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Item{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
