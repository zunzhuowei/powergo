package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"httpServer/controllers"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

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
		beego.NSNamespace("/members",
			beego.NSInclude(
				&controllers.MembersController{},
			),
		),
	)
	beego.AddNamespace(ns)

}
