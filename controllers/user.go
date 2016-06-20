package controllers

import (
	"forfish/models"
	"forfish/services"
	"github.com/astaxie/beego"
	"time"
)

// UserController 主页控制器
type UserController struct {
	beego.Controller
}

// Get 注册页面
func (u *UserController) Get() {
	u.TplName = "user/register.html"
}

// Post 注册请求
//  account:账号
//  password:密码
func (u *UserController) Post() {
	defer func() {
		u.ServeJSON()
	}()
	user := models.User{
		Account:    u.GetString("account"),
		Password:   u.GetString("password"),
		CreateTime: time.Now(),
	}
	server := services.DefaultServer
	id, err := server.Add(user)
	if err != nil {
		u.Data["json"] = map[string]string{"error": err.Error()}
		return
	}
	u.Data["json"] = id
}