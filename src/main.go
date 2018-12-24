package main

import (
	"github.com/astaxie/beego"
	_ "httpServer/routers"
)

func main() {
	// tests goroutine and channel
	//tests.TestChannel()
	//tests.TestChannel2()

	beego.Run()
}
