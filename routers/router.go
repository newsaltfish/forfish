package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/index", &controllers.IndexController{}, "get:LoginView;post:Login")
	//用户信息相关
	beego.Router("/user/register", &controllers.UserController{})
	beego.Router("/test", &controllers.TestController{})
}
