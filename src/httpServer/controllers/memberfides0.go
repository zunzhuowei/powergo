package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"httpServer/models"
	"strconv"
)

type Memberfides0Controller struct {
	beego.Controller
}

func (controller *Memberfides0Controller) GetOneMemberfides() {
	var mem = map[string]interface{}{}

	var mid = controller.Ctx.Input.Param(":mid")
	v, _ := strconv.Atoi(mid)
	mem["mid"] = v
	mem["name"] = "zhangsan"
	mem["sex"] = 1

	city := controller.GetString("city")
	mem["city"] = city

	controller.Data["json"] = mem
	controller.ServeJSON()

}

func (controller *Memberfides0Controller) GetMemberfidesList() {
	start, _ := strconv.Atoi(controller.Ctx.Input.Param(":start"))
	end, _ := strconv.Atoi(controller.Ctx.Input.Param(":end"))
	memberfide := models.Memberfides0{}
	res := (&memberfide).GetMemberfidesList(start, end)
	beego.Debug("Memberfides0Controller GetMemberfidesList res:", res)
	controller.Data["json"] = res
	fmt.Println(start, end)
	beego.Debug("Memberfides0Controller GetMemberfidesList start end:", start, end)

	controller.ServeJSON()
}
