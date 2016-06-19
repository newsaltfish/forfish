package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id         int
	Account    string
	NickName   string
	Password   string
	CreateTime time.Time
}
