package controllers

import "github.com/astaxie/beego"

// IndexController 主页控制器
type IndexController struct {
	beego.Controller
}

// MainController 主页
type MainController struct {
	beego.Controller
}

// Get 进入登录页面
func (c *MainController) Get() {
	c.TplName = "index/index.html"
}

// Post 进入登录页面
func (c *MainController) Post() {
	c.Data["data"] = map[string]interface{}{"code": 1000, "msg": "嘿嘿嘿嘿", "请求参数": c.Ctx.Request.Form}
	c.ServeJSON()
}


// LoginView 进入登录页面
func (c *IndexController) LoginView() {
	c.TplName = "index/index.html"
}

// PicCode 图形验证码
func (c *IndexController) PicCode() {
}

// Login 登录请求
func (c *IndexController) Login() {
	// account := c.GetString("account")
	// password := c.GetString("password")
	// code := c.GetString("piccode")
}
