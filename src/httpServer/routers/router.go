package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
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

	if beego.BConfig.WebConfig.Session.SessionOn {
		loginFilter()
	} else {
		sessionKeyFilter()
	}
	namespaceMapping()

	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/user", &controllers.MainController{}, "get:UserList")
	beego.Router("/login", &controllers.MainController{}, "get:UserList")
	beego.Router("/api/user", &controllers.MainController{}, "get:UserListApi")
	beego.Router("/api/user/:id([0-9]+)", &controllers.MainController{}, "get:UserListApi2")
	beego.Router("/api/user2/:id:int", &controllers.MainController{}, "get:UserListApi2")
	beego.Router("/api/fides/:mid:int", &controllers.Memberfides0Controller{}, "get:GetOneMemberfides")
	beego.Router("/api/fides/list/:start:int/:end:int", &controllers.Memberfides0Controller{}, "get:GetMemberfidesList")

}

// 命名空间映射
func namespaceMapping() {
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

// 如果使用session 模块，则校验session 中的uid
func loginFilter() *beego.App {
	return beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
		uid := ctx.Input.Session("uid")

		_, ok := uid.(int)
		beego.Debug("----------", ok)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
	})
}

// 如果采用无状态的,则 校验session key
func sessionKeyFilter() *beego.App {
	return beego.InsertFilter("/*", beego.BeforeRouter, func(context *context.Context) {
		values := context.Request.Form
		for k, v := range values {
			beego.Debug("sessionKeyFilter222222222222 key value-----::", k, v)
		}

		for k, v := range context.Input.Params() {
			beego.Debug("sessionKeyFilter key value-----::", k, v)
		}
	})
}
