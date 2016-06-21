package services

import (
	"github.com/astaxie/beego/orm"
)

// BaseServer 基础server
type BaseServer struct {
	Orm orm.Ormer
}

func init() {

}
