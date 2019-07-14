package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["mini-market-go/controllers:AccountController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:AccountController"],
        beego.ControllerComments{
            Method: "GetAccount",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:AccountController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:AccountController"],
        beego.ControllerComments{
            Method: "Authenticate",
            Router: `/authenticate`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:ImageController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:ImageController"],
        beego.ControllerComments{
            Method: "GetByItemId",
            Router: `/get-by-item-id/:itemId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:ImageController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:ImageController"],
        beego.ControllerComments{
            Method: "Upload",
            Router: `/upload`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:ItemController"],
        beego.ControllerComments{
            Method: "Search",
            Router: `/get-by-customer`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:TypeController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:TypeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:TypeController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:TypeController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:TypeController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:TypeController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:TypeController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:TypeController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["mini-market-go/controllers:TypeController"] = append(beego.GlobalControllerRouter["mini-market-go/controllers:TypeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
