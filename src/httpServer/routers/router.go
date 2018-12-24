package routers

import (
	"github.com/astaxie/beego"
	"httpServer/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
