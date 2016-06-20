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
	Conns = make(map[string]*websocket.Conn)
	// GameConns 位置信息
	GameConns = make(map[*websocket.Conn]string)
	// GmsChannel 游戏位置信息
	GmsChannel = make(chan string, 10)
	locker     sync.RWMutex
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

// GameEcho 画图信息
func GameEcho(ws *websocket.Conn) {
	beego.Debug("new user", ws)
	defer func() {
		beego.Debug("leave")
	}()
	player := ""
	if err := websocket.Message.Receive(ws, &player); err != nil {
		beego.Debug("连接错误" + err.Error())
	}
	if player != "" {
		GameConns[ws] = "player"
	} else {
		GameConns[ws] = "guess"
	}
	for {
		reply := ""
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			beego.Debug("连接错误" + err.Error())
			break
		}

		// fmt.Println("位置信息:" + reply)
		GmsChannel <- reply
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
		case gms := <-GmsChannel:
			for k, _ := range GameConns {

				beego.Debug(websocket.Message.Send(k, gms))
			}
		}
	}
}
