package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"mini-market-go/models"
)

//  ImageController operations for Image
type ImageController struct {
	beego.Controller
}

// @router /upload [post]
func (c *AccountController) Upload() {

	fileHeaders, err := c.GetFiles("myfiles")
	if err != nil {
		glog.Errorf("Error get file: %s", err.Error())
	} else {
		for _, file := range fileHeaders {
			glog.Infof("FILE: %v", file.Filename)
			models.SaveFile(file)
		}
	}
	c.ServeJSON()
}
