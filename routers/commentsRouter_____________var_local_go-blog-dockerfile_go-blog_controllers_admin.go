package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["go-blog/controllers/admin:AdController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:AdController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/ad",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:AdController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:AdController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/ad",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:AdController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:AdController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/ad/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:AdController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:AdController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/ad/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:AdController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:AdController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/ad/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:CronController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:CronController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "cron",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:CronController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:CronController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "cron/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:CronController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:CronController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "cron/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:CronController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:CronController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "cron/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:CronController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:CronController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "cron/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:CustomerController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:CustomerController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/customer",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:CustomerController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:CustomerController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/customer/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:CustomerController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:CustomerController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/customer/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:CustomerController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:CustomerController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/customer/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:CustomerController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:CustomerController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/customer/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/file/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/file/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/file/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/file/add",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/file/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:FileController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/files",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/link",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/link/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/link/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/link/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/link/add",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:LinkController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/link/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/menu",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/menu/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/menu/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/menu/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Add",
            Router: "/menu/add",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"] = append(beego.GlobalControllerRouter["go-blog/controllers/admin:MenuController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/menu/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
