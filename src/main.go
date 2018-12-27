package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "httpServer/routers"
	"io/ioutil"
	"runtime"
)

func init() {
	/*err := orm.RegisterDataBase("default", "mysql",
		username+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		fmt.Println(err)
	}*/

}

var (
	username string
	password string
	host     string
	port     string
	dbname   string
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

	var serverJson map[string]interface{}

	err = json.Unmarshal(data, &serverJson)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	connectDb(serverJson)
	//beego.SetLogger("file", str)
	//beego.BeeLogger.DelLogger("console")
	//beego.SetLevel(beego.LevelInformational)
	//beego.SetLevel(beego.LevelError)
	beego.SetLogFuncCall(true)

	//beego.SetStaticPath("/static","/src/httpServer/static")

	beego.Run()
}

func connectDb(serverJson map[string]interface{}) {
	username = string(serverJson["mysql.username"].(string))
	password = string(serverJson["mysql.password"].(string))
	host = string(serverJson["mysql.host"].(string))
	port = string(serverJson["mysql.port"].(string))
	dbname = string(serverJson["mysql.dbName"].(string))

	//err := orm.RegisterDataBase("default", "mysql",
	//	username+":"+password+"@tcp("+host+":"+port+")/"+dbname)

	maxIdle := 30
	maxConn := 30
	err := orm.RegisterDataBase("default", "mysql",
		username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8",
		maxIdle, maxConn)

	if err != nil {
		fmt.Println(err)
	}
}
