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
	Id      int64  `orm:"column(id);auto" json:"id"`
	Name    string `orm:"column(name);size(255);null" json:"name"`
	Content string `orm:"column(content);size(255);null" json:"content"`
	Item    *Item  `orm:"column(item);rel(fk);null;on_delete(set_null)" json:"item"`
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

func GetImageByItemId(itemId int) (images []*Image, err error) {
	o := orm.NewOrm()
	if err = o.QueryTable(new(Image)).Filter("item", itemId).One(&images); err == nil {
		for _, im := range images {
			im.Item = nil
		}
		return images, nil
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

func SaveFile(file *multipart.FileHeader, itemId int64) error {
	splits := strings.Split(file.Filename, ".")
	ext := (map[bool]string{true: splits[len(splits)-1], false: "png"})[len(splits) > 1]
	content, err := file.Open()
	defer content.Close()
	if err != nil {
		glog.Errorf("Open file error: %s", err.Error())
		return err
	}
	bytes := make([]byte, 100000000)
	_, err = content.Read(bytes)
	if err != nil {
		glog.Errorf("Read content to bytes error: %s", err.Error())
		return err
	}
	newName := fmt.Sprintf("%s.%s", SHA1(bytes), ext)
	glog.Infof("new name: %s", newName)
	// Save
	dst, err := os.Create("/home/dunghv/images_upload/" + newName)
	defer dst.Close()
	if err != nil {
		glog.Errorf("Create directory failed: %s", err.Error())
		return err
	}
	if _, err := dst.Write(bytes); err != nil {
		glog.Errorf("Write file to directory err: %s", err.Error())
		return err
	}
	// glog.Infof("Write file success: %s", file.Filename)
	// save to DB
	if isExistFile(newName, itemId) {
		glog.Warningf("Item id=%v already exist file name: %s", itemId, newName)
		return nil
	}
	img := Image{
		Id:      0,
		Name:    newName,
		Content: newName,
		Item:    &Item{Id: itemId},
	}
	o := orm.NewOrm()
	o.QueryTable(new(Image))
	_, err = o.Insert(&img)
	if err != nil {
		glog.Errorf("save image to DB err: %s", err.Error())
		return err
	}
	glog.Infof("Save image to DB success: %s", img.Name)
	return nil
}

func SHA1(bytes []byte) string {
	h := sha1.New()
	h.Write(bytes)
	sum := h.Sum(nil)
	return fmt.Sprintf("%x", sum)
}

func isExistFile(fileName string, itemId int64) bool {
	o := orm.NewOrm()
	return o.QueryTable(new(Image)).Filter("name", fileName).Filter("item", itemId).Exist()
}
