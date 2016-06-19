package services

import (
	"github.com/astaxie/beego/orm"
	"hello/models"
	_ "hello/utils"
)

var (
	DefaultServer UserService
)

// UserService 用户相关操作
type UserService struct {
	Orm orm.Ormer
}

func init() {
	var o = orm.NewOrm()
	DefaultServer.Orm = o
}
func (u *UserService) Add(user models.User) (int, error) {
	n, err := u.Orm.Insert(&user)
	return int(n), err
}
