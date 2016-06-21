package controllers

import (
	"forfish/utils"

	"github.com/astaxie/beego"
)

// TestController 测试
type TestController struct {
	beego.Controller
}

// Post 进入登录页面
func (c *TestController) Post() {
	c.Data["json"] = map[string]interface{}{"code": 1000, "msg": "嘿嘿嘿嘿", "请求参数": c.Ctx.Request.Form}
	c.ServeJSON()
}

// Get 进入游戏页面
func (c *TestController) Get() {
	userName := c.GetSession("user")     //判断是否是已进入过游戏的用户
	if _, ok := userName.(string); !ok { //已断开连接 需要重连

	}
	// c.Data["user"] = utils.Game.Users
	c.TplName = "guess/index.html"
}

// LoginView 进入登录页面
func (c *TestController) LoginView() {
	c.TplName = "guess/login.html"
}

// Login 登录
func (c *TestController) Login() {
	uname := c.GetString("name")
	defer func() {
		c.ServeJSON()
	}()
	if uname == "" {
		c.Data["json"] = utils.ReturnFailure(1)
		return
	}
	c.SetSession("user", uname)
	c.Data["json"] = utils.ReturnSuccess(nil)
}
