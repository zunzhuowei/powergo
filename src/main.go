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
	connectDb()
}

func main() {

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

	beego.SetLogger("file", string(data))
	//beego.BeeLogger.DelLogger("console")
	//beego.SetLevel(beego.LevelInformational)
	//beego.SetLevel(beego.LevelError)
	beego.SetLogFuncCall(true)

	beego.Run()
}

type dbModel struct {
	username string
	password string
	host     string
	port     string
	dbname   string
}

func connectDb() {
	data, err := ioutil.ReadFile("conf/db.json")
	var dbJson map[string]interface{}

	err = json.Unmarshal(data, &dbJson)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for dbName, v := range dbJson {
		m := v.(map[string]interface{})
		username := m["username"].(string)
		password := m["password"].(string)
		host := m["host"].(string)
		port := m["port"].(string)
		dbname := m["dbName"].(string)
		maxIdle := m["maxIdle"].(float64)
		maxConn := m["maxConn"].(float64)
		err = orm.RegisterDataBase(dbName, "mysql",
			username+":"+password+"@tcp("+host+":"+port+")/"+dbname+"?charset=utf8",
			int(maxIdle), int(maxConn))
		beego.Debug(dbName, v)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
