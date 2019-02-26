package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["mini-market-go/controllers:AccountController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:AccountController"],
		beego.ControllerComments{
			Method:           "Authenticate",
			Router:           `/authenticate`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
