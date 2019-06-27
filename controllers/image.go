package controllers

import (
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"mini-market-go/models"
	"net/http"
)

//  ImageController operations for Image
type ImageController struct {
	beego.Controller
}

// @Param itemId query int true "ID of Item has an image"
// @router /upload [post]
func (c *ImageController) Upload() {
	itemId, err := c.GetInt64("itemId")
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, "image.upload.error.missing-itemId")
	}
	fileHeader, err := c.GetFiles("file")
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, "image.upload.error.get-file-error")
		glog.Errorf("Error get file: %s", err.Error())
	}
	glog.Infof("files: %v", fileHeader)
	for _, file := range fileHeader {
		err = models.SaveFile(file, itemId)
	}
	if err != nil {
		c.CustomAbort(http.StatusBadRequest, "image.upload.error.save-file-error")
	}
	c.Data["json"] = "upload-success"
	c.ServeJSON()
}
