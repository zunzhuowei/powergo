package routers

import (
	"github.com/astaxie/beego"
	"httpServer/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/user", &controllers.MainController{}, "get:UserList")
	beego.Router("/api/user", &controllers.MainController{}, "get:UserListApi")
	beego.Router("/api/user/:id([0-9]+)", &controllers.MainController{}, "get:UserListApi2")
	beego.Router("/api/user2/:id:int", &controllers.MainController{}, "get:UserListApi2")

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)

}
