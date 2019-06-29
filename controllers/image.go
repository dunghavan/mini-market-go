package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/golang/glog"
	"mini-market-go/models"
	"net/http"
	"strconv"
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

// @Param itemId query int true "ID of Item want to get images"
// @router /get-by-item-id/:itemId [get]
func (c *ImageController) GetByItemId() {
	idStr := c.Ctx.Input.Param(":itemId")
	itemId, err := strconv.Atoi(idStr)
	if err != nil || itemId == 0 {
		c.CustomAbort(http.StatusBadRequest, "image.get-by-item-id.error.invalid-body")
	}
	images, err := models.GetImageByItemId(itemId)
	if err != nil {
		glog.Errorf("Get images by itemId=%v err: %s", err.Error())
		if err == orm.ErrNoRows {
			images = make([]*models.Image, 0)
		} else {
			c.CustomAbort(http.StatusBadRequest, "image.get-by-item-id.error.query-db-error")
		}
	}
	c.Data["json"] = images
	c.ServeJSON()
}
