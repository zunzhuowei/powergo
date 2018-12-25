package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "httpServer/routers"
	"io/ioutil"
	"runtime"
)

func main() {
	// tests goroutine and channel
	//tests.TestChannel()
	//tests.TestChannel2()

	httpServer()
}

func httpServer() {
	defer func() {
		if err := recover(); err != nil {
			beego.Alert("recover httpServer err : ", err)
		}
	}()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.BConfig.Listen.ServerTimeOut = 10

	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		return
	}

	if runtime.GOOS != "windows" {
		beego.BeeLogger.DelLogger("console") //删除控制台输出
	}

	//log
	//beego.SetLogger("file", `{"filename":"bin/beegoLog/test.log"}`)
	var str string = string(data)

	fmt.Println(str)

	//beego.SetLogger("file", str)
	//beego.BeeLogger.DelLogger("console")
	//beego.SetLevel(beego.LevelInformational)
	//beego.SetLevel(beego.LevelError)
	beego.SetLogFuncCall(true)

	//beego.SetStaticPath("/static","/src/httpServer/static")

	beego.Run()
}
