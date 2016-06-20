package utils

import (
	// _ "demoyb/routers"
	// "fmt"
	"fmt"

	"golang.org/x/net/websocket"
	// "github.com/astaxie/beego"
)

func init() {
	// beego.Run()

	// http.HandleFunc("/123", websocket.Handler(Echo))
}
func Echo(ws *websocket.Conn) {
	var err error
	fmt.Println("開始")
	for {
		reply := ""
		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("连接错误" + err.Error())
			break
		}
		fmt.Println("客户端发送:" + reply)
		if err = websocket.Message.Send(ws, "服务器发送的字段:嘿嘿嘿"); err != nil {
			break
		}

	}
}
