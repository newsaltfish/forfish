package main

import (
	"fmt"
	_ "forfish/routers"
	"forfish/utils"
	"log"
	"net/http"

	"github.com/astaxie/beego"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/hello", websocket.Handler(utils.Echo))
	beego.Handler("/", h, options)

	log.Fatal("starting")
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Println("结束")
	beego.Run()
}
