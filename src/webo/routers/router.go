package routers

import (
	"github.com/astaxie/beego"
	"webo/controllers"
)

func init() {
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/service", &controllers.ServiceController{})
	beego.Router("/item/get/?:id", &controllers.ItemController{}, "*:Get")
	beego.Router("/item/list/:hi:string", &controllers.ItemController{}, "*:List")
	beego.Router("/item/add", &controllers.ItemController{}, "*:Add")
	beego.Router("/item/update", &controllers.ItemController{}, "*:Update")
	beego.Router("/item/delete", &controllers.ItemController{}, "*:Delete")
	beego.Router("/ui/add/:hi:string", &controllers.UiController{}, "*:Add")
}
