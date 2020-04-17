package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/controllers:UserController"] = append(beego.GlobalControllerRouter["crud_with_beego/Golang-POC-BEEGO/crud-poc-beego/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
