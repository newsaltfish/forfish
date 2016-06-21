package main

import (
	_ "forfish/routers"
	"forfish/utils"

	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
)

func main() {
	beego.Handler("/hello", websocket.Handler(utils.Echo))
	beego.Handler("/game", websocket.Handler(utils.GameEcho))
	beego.Run()
}
