package routers

import (
	"github.com/astaxie/beego"
	"httpServer/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/user", &controllers.MainController{}, "get:UserList")
	beego.Router("/api/user", &controllers.MainController{}, "get:UserListApi")
}
