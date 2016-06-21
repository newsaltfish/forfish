package utils

import (
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
	GameConns = make(map[string]*websocket.Conn)
	// GmsChannel 游戏位置信息
	GmsChannel = make(chan string, 10)
	locker     sync.RWMutex
)

// Echo 建立连接
func Echo(ws *websocket.Conn) {
	user := ""
	if err := websocket.Message.Receive(ws, &user); err != nil {
		beego.Debug("连接错误" + err.Error())
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
			beego.Debug("连接错误" + err.Error())
			break
		}
		MsgChannel <- user + ":" + reply
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
	locker.Lock()
	GameConns[player] = ws
	locker.Unlock()
	for {
		reply := ""
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			beego.Debug("连接错误" + err.Error())
			locker.Lock()
			delete(GameConns, player)
			locker.Unlock()
			break
		}
		GmsChannel <- reply
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
			for _, v := range GameConns {
				websocket.Message.Send(v, gms)
			}
		}
	}
}
