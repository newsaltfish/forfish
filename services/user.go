package services

import (
	"forfish/models"
	_ "forfish/utils"
	"github.com/astaxie/beego/orm"
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
