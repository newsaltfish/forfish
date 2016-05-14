package models

import "time"

type User struct {
	Id         int
	Account    string
	NickName   string
	LoginTime  time.Time
	CreateTime time.Time
}
