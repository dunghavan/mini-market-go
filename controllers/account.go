package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"mini-market-go/models"
	"mini-market-go/service"
	"net/http"
)

type AccountController struct {
	beego.Controller
}

// @Param	body		body 	models.LoginVM	true 	"authenticate login model"
// @router /authenticate [post]
func (c *AccountController) Authenticate() {
	var l models.LoginVM
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &l); err == nil {
		fmt.Printf("%v", l.Email)
		if res, status := service.GetFacebookInfo(l.FBToken); status == http.StatusOK {
			var fbUser models.FaceBookUser
			if err := json.Unmarshal(res, &fbUser); err == nil {
				fmt.Printf("%v", fbUser)
			} else {
				glog.Errorf("Parse fbUser error: %s", err.Error())
			}
		} else {
			glog.Errorf("Get facebook user info error status=%v", status)
			c.CustomAbort(http.StatusBadRequest, "account.login.error.facebookusernotfound")
		}
	} else {
		c.CustomAbort(http.StatusBadRequest, "account.login.error.invaliddata")
	}
	c.ServeJSON()
}
