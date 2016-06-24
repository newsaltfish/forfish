package services

import (
	"github.com/astaxie/beego/orm"
)

// BaseService 基础service
type BaseService struct {
	Orm orm.Ormer
}

// BaseServer 基础server
var BaseServer BaseService

func init() {
	BaseServer.Orm = orm.NewOrm()
}
