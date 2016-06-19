package models

import "time"

type User struct {
	Id         int
	Account    string
	NickName   string
	CreateTime time.Time
}
