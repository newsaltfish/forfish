package routers

import (
	"hello/controllers"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/index", &controllers.IndexController{}, "get:LoginView;post:Login")
	beego.Router("/test", &controllers.TestController{})
}
