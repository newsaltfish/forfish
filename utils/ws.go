package utils

import (
	"fmt"
	"sync"

	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
)

func init() {
	go MsgSend()
}

var (
	// MsgChannel 信息
	MsgChannel = make(chan string, 10)
	// Conns 连接
	Conns  = make(map[string]*websocket.Conn)
	locker sync.RWMutex
)

// Echo 建立连接
func Echo(ws *websocket.Conn) {
	user := ""
	if err := websocket.Message.Receive(ws, &user); err != nil {
		fmt.Println("连接错误" + err.Error())
	} else {
		locker.Lock()
		Conns[user] = ws
		locker.Unlock()
	}
	beego.Debug(Conns)
	defer func() {
		locker.Lock()
		delete(Conns, user)
		locker.Unlock()
	}()
	for {
		reply := ""
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("连接错误" + err.Error())
			break
		}
		fmt.Println("客户端发送:" + reply)
		MsgChannel <- user + ":" + reply
		// if err := websocket.Message.Send(ws, "服务器发送的字段:嘿嘿嘿"); err != nil {
		// 	break
		// }
	}
}

// MsgSend 发送信息
func MsgSend() {
	for {
		select {
		case ms := <-MsgChannel:
			for _, v := range Conns {
				websocket.Message.Send(v, ms)
			}
		}
	}
}
