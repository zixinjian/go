package routers

import (
	"github.com/astaxie/beego"
	"webo/controllers"
)

func init() {
	//框架服务
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/main", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/service", &controllers.ServiceController{})

	//基本资源服务
	beego.Router("/item/get/?:id", &controllers.ItemController{}, "*:Get")
	beego.Router("/item/list/:hi:string", &controllers.ItemController{}, "*:List")
	beego.Router("/item/add/:hi:string", &controllers.ItemController{}, "*:Add")
	beego.Router("/item/update/:hi:string", &controllers.ItemController{}, "*:Update")
	beego.Router("/item/delete/:hi:string", &controllers.ItemController{}, "*:Delete")
	beego.Router("/item/upload/:hi:string", &controllers.ItemController{}, "*:Upload")
	beego.Router("/item/autocomplete/:hi:string", &controllers.ItemController{}, "*:Autocomplete")

	//基本资源页面
	beego.Router("/ui/add/:hi:string", &controllers.UiController{}, "*:Add")
	beego.Router("/ui/list/:hi:string", &controllers.UiController{}, "*:List")
	beego.Router("/ui/update/:hi:string", &controllers.UiController{}, "*:Update")

	//用户管理
	beego.Router("/ui/user/list", &controllers.UserController{}, "*:UiList")
	beego.Router("/item/update/user", &controllers.UserController{}, "*:Update")

	//产品管理
	beego.Router("/ui/product/add", &controllers.ProductController{}, "*:UiAdd")
	beego.Router("/ui/product/list", &controllers.ProductController{}, "*:UiList")
	beego.Router("/ui/product/update", &controllers.ProductController{}, "*:UiUpdate")
	beego.Router("/item/product/list", &controllers.ProductController{}, "*:List")

	//采购管理
	beego.Router("/ui/purchase/mycreate", &controllers.PurchaseController{}, "*:UiMyCreate")
	beego.Router("/ui/purchase/curlist", &controllers.PurchaseController{}, "*:UiCurList")
	beego.Router("/ui/purchase/history", &controllers.PurchaseController{}, "*:UiHistoryList")
	beego.Router("/ui/purchase/add", &controllers.PurchaseController{}, "*:UiAdd")
	beego.Router("/ui/purchase/update", &controllers.PurchaseController{}, "*:UiUpdate")
	beego.Router("/ui/purchase/userupdate", &controllers.PurchaseController{}, "*:UiUserUpdate")
	beego.Router("/ui/purchase/show", &controllers.PurchaseController{}, "*:UiHistoryUpdate")
	beego.Router("/item/list/purchase", &controllers.PurchaseController{}, "*:List")

	//分析
	beego.Router("/ui/purchase/priceAnalyze", &controllers.PurchaseController{}, "*:PriceAnalyze")
	beego.Router("/ui/expense/List", &controllers.PurchaseController{}, "*:ExpenseList")

	//出差管理
	beego.Router("/travel", &controllers.MainController{}, "*:Travel")
}
