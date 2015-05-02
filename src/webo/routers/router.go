package routers

import (
	"webo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.LoginController{})
    beego.Router("/login", &controllers.LoginController{})
	beego.Router("/user", &controllers.UserController{})
//	beego.Router("/service", &controllers.ServiceController{})
}
