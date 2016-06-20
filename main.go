package main

import (
	"fmt"
	_ "forfish/routers"
	"forfish/utils"

	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
)

func main() {
	//		http.Handle("/hello", websocket.Handler(utils.Echo))
	beego.Handler("/hello", websocket.Handler(utils.Echo))
	beego.Handler("/game", websocket.Handler(utils.GameEcho))

	//	log.Fatal("starting")
	//	if err := http.ListenAndServe(":1234", nil); err != nil {
	//		log.Fatal("ListenAndServe:", err)
	//	}
	fmt.Println("结束")
	beego.Run()
}
