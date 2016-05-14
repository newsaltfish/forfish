package controllers

import "github.com/astaxie/beego"

// IndexController 主页
type IndexController struct {
	beego.Controller
}

// Get get请求
func (c *IndexController) Get() {
	c.TplName = "index/index.html"
}