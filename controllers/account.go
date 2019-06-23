package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/golang/glog"
	"mini-market-go/models"
	"mini-market-go/security"
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
		glog.Infof("User login with fbId=%s", l.FBId)
		if res, status := service.GetFacebookInfo(l.FBToken); status == http.StatusOK {
			var fbUser models.FaceBookUser
			if err = json.Unmarshal(res, &fbUser); err == nil {
				if u, err := models.CreateOrUpdate(&fbUser); err == nil {
					token := security.CreateUserToken(u, "")
					c.Data["json"] = fmt.Sprintf(`{"id_token": "%s"}`, token)
					c.Ctx.Output.Header("Authorization", fmt.Sprintf("Bearer %s", token))
				} else {
					c.CustomAbort(http.StatusBadRequest, "account.login.error.insertuserfailed")
				}
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

// @router / [get]
func (c *AccountController) GetAccount() {
	userId, err := security.GetCurrentUserId(&c.Controller)
	if err != nil {
		glog.Infof("get current user id err: %s", err.Error())
		c.CustomAbort(http.StatusBadRequest, "accounts.get.error.get-current-user-error")
	}
	if user, err := models.GetUserById(userId); err != nil {
		glog.Error("Get user by id err: %s", err.Error())
		c.CustomAbort(http.StatusBadRequest, "accounts.get.error.get-user-by-id-error")
	} else {
		c.Data["json"] = user
	}
	c.ServeJSON()
}
