package services

import (
	"forfish/models"
	"forfish/utils"

	"github.com/astaxie/beego/orm"
)

var defaultOrm orm.Ormer

func init() {
	defaultOrm = orm.NewOrm()
}

// UserService 用户相关操作
type UserService struct {
	BaseService
}

// Add 新增用户
func (u *UserService) Add(user models.User) (int, error) {
	n, err := BaseServer.Orm.Insert(&user)
	return int(n), err
}

// Login 登录
func (u *UserService) Login(account, password string) bool {
	sql := `select count(*) from user where account=? and password=?`
	var num int
	err := BaseServer.Orm.Raw(sql, account, utils.Md5(password)).QueryRow(&num)
	if err != nil || num < 1 {
		return false
	}
	return true
}
