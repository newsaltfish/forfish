package controllers

import "github.com/astaxie/beego"

// TestController 测试
type TestController struct {
	beego.Controller
}

// Post 进入登录页面
func (c *TestController) Post() {
	c.Data["json"] = map[string]interface{}{"code": 1000, "msg": "嘿嘿嘿嘿", "请求参数": c.Ctx.Request.Form}
	c.ServeJSON()
}
