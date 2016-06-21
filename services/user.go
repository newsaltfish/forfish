package services

import (
	"forfish/models"

	"github.com/astaxie/beego/orm"
)

var defaultOrm orm.Ormer

func init() {
	defaultOrm = orm.NewOrm()
}

// UserService 用户相关操作
type UserService struct {
	BaseServer
}

// Add 新增用户
func (u *UserService) Add(user models.User) (int, error) {
	n, err := u.Orm.Insert(&user)
	return int(n), err
}
