package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/golang/glog"
	"mini-market-go/controllers"
	"mini-market-go/security"
	"net/http"
	"time"
)

var FilterUser = func(ctx *context.Context) {
	if err := security.VerifyRequest(ctx); err != nil {
		ctx.Abort(http.StatusUnauthorized, "unauthorized")
	}
}

func init() {
	crossDomain()
	router()
}

func crossDomain() {
	var trustDomain string
	if beego.BConfig.RunMode == beego.PROD {
		trustDomain = beego.AppConfig.String("TrustDomain")
		glog.Infof("Allow trust domain %s", trustDomain)
	} else {
		trustDomain = "*"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{trustDomain},
		AllowCredentials: true,
		AllowMethods:     []string{"OPTION", "GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"DNT", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Range", "Authorization"},
		ExposeHeaders:    []string{"X-Total-Count", "Content-Length", "Content-Range, Authorization"},
		MaxAge:           24 * time.Hour,
	}))
}

func router() {
	beego.InsertFilter("/core/*", beego.BeforeRouter, FilterUser)

	nsapi := beego.NewNamespace("/core/v1",
		beego.NSNamespace("/accounts",
			beego.NSInclude(
				&controllers.AccountController{
					Controller: beego.Controller{},
				},
			),
		),
		beego.NSNamespace("/images",
			beego.NSInclude(
				&controllers.ImageController{
					Controller: beego.Controller{},
				},
			),
		),
		beego.NSNamespace("/items",
			beego.NSInclude(
				&controllers.ItemController{
					Controller: beego.Controller{},
				},
			),
		),
	)
	beego.AddNamespace(nsapi)
}
