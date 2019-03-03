package models

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
	"mime/multipart"
	"os"
	"strings"
)

type Image struct {
	Id      int64  `orm:"auto"`
	Name    string `orm:"size(128);null" json:"name"`
	Content string `orm:"size(128);null" json:"content"`
}

func init() {
	orm.RegisterModel(new(Image))
}

// AddImage insert a new Image into database and returns
// last inserted Id on success.
func AddImage(m *Image) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetImageById retrieves Image by Id. Returns error if
// Id doesn't exist
func GetImageById(id int64) (v *Image, err error) {
	o := orm.NewOrm()
	v = &Image{Id: id}
	if err = o.QueryTable(new(Image)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// DeleteImage deletes Image by Id and returns error if
// the record to be deleted doesn't exist
func DeleteImage(id int64) (err error) {
	o := orm.NewOrm()
	v := Image{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Image{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func SaveFile(file *multipart.FileHeader) {
	splits := strings.Split(file.Filename, ".")
	ext := (map[bool]string{true: splits[len(splits)-1], false: "png"})[len(splits) > 1]
	content, err := file.Open()
	defer content.Close()
	if err != nil {
		glog.Errorf("Open file error: %s", err.Error())
		return
	}
	bytes := make([]byte, 100000000)
	_, err = content.Read(bytes)
	if err != nil {
		glog.Errorf("Read content to bytes error: %s", err.Error())
		return
	}
	newName := fmt.Sprintf("%s.%s", SHA1(bytes), ext)
	glog.Infof("new name: %s", newName)
	// Save
	dst, err := os.Create("/home/dunghv/sanat/" + newName)
	defer dst.Close()
	if err != nil {
		glog.Errorf("Create directory failed: %s", err.Error())
		return
	}
	if _, err := dst.Write(bytes); err != nil {
		glog.Errorf("Write file to directory err: %s", err.Error())
		return
	}
	glog.Infof("Write file success: %s", file.Filename)
	// TODO save to DB
}

func SHA1(bytes []byte) string {
	h := sha1.New()
	h.Write(bytes)
	sum := h.Sum(nil)
	return fmt.Sprintf("%x", sum)
}
