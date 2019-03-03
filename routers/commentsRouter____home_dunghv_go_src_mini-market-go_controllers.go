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

	beego.GlobalControllerRouter["mini-market-go/controllers:AccountController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:AccountController"],
		beego.ControllerComments{
			Method:           "Upload",
			Router:           `/upload`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mini-market-go/controllers:UserController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mini-market-go/controllers:UserController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mini-market-go/controllers:UserController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mini-market-go/controllers:UserController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["mini-market-go/controllers:UserController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
