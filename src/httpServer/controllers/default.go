package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (This *MainController) Get() {
	This.Data["Website"] = "beego.me"
	This.Data["Email"] = "astaxie@gmail.com"
	This.TplName = "index.tpl"
}

func (This *MainController) UserList() {
	This.Data["name"] = "张三"
	This.Data["tel"] = "1234564323"
	This.TplName = "users/user_list.tpl"
}

func (This *MainController) UserListApi() {
	This.Data["name"] = "张三"
	This.Data["tel"] = "1234564323"
	This.Ctx.WriteString("hello")
}

func (This *MainController) UserListApi2() {
	var id string = This.Ctx.Input.Param(":id")
	fmt.Println(id)
	This.Data["name"] = "张三"
	This.Data["tel"] = "1234564323"
	This.Ctx.WriteString("hello")
}
