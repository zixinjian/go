package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"webo/models/rpc"
	"webo/models/userMgr"
)

type UserController struct {
	beego.Controller
}

func (self *UserController) Get() {
	users, _ := userMgr.GetUserList()
	tr := rpc.TableResult{}
	tr.Rows = users
	tr.Total = len(users)
	self.Data["json"] = &tr
	self.ServeJson()
}

func (self *UserController) Post() {
	fmt.Println(self.Ctx.Input.RequestBody)
	fmt.Println(self.Input())
	self.ParseForm()
	ret := rpc.JsonResult{"success", "id"}
	self.Data["json"] = &ret
	self.ServeJson()
}
