package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"httpServer"
	"io/ioutil"
	"server"
)

func main() {
	data, err := ioutil.ReadFile("conf/powergo.json")
	var powergoMap map[string]interface{}

	err = json.Unmarshal(data, &powergoMap)
	if err != nil {
		fmt.Println("conf/powergo.json -------::", err.Error())
		return
	}

	enable := powergoMap["enable"].(map[string]interface{})

	enableHttpServer := enable["httpServer"].(bool)
	enableGameServer := enable["gameServer"].(bool)

	if enableGameServer && enableHttpServer {
		beego.Warn("enable http server and game server !")
		go httpServer.HttpServer()
		server.GameServer()
	} else {
		if enableGameServer {
			beego.Warn("only enable game server !")
			server.GameServer()
		} else {
			if enableHttpServer {
				beego.Warn("only enable http server !")
				httpServer.HttpServer()
			} else {
				beego.Warn("no enable any server !")
			}
		}
	}

}
