package utils

import (
	"sync"

	"golang.org/x/net/websocket"
)

type gamestatus int //游戏状态

const (
	_ gamestatus = iota
	// GameWait 等待进行
	GameWait
	// GameStart 游戏开始
	GameStart
	//GameEnd 游戏结束
	GameEnd
)

// Game 游戏类型
type Game struct {
	Users []User
	Canvas
	Status gamestatus //游戏状态
}

// User 用户
type User struct {
	Name string
	Role string
}

// Canvas 画布
type Canvas struct {
	Points []Point
	Color
}

// Point 画布的像素点
type Point struct {
	X int
	Y int
	Color
}

// Color 颜色类型
type Color struct {
	Color int
	Alpha int
}

var (
	// MsgChannel 信息
	MsgChannel = make(chan string, 10)
	// Conns 连接
	Conns = make(map[string]*websocket.Conn)
	// GameConns 位置信息
	// GameConns = make(map[string]*websocket.Conn)
	// GameConns 位置信息
	GameConns []*websocket.Conn
	// GmsChannel 游戏位置信息
	GmsChannel = make(chan string, 10)
	locker     sync.RWMutex //锁
)
