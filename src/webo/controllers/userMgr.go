package controllers

import (
    "github.com/astaxie/beego"
    "webo/models/userMgr"
    "webo/models/rpc"
    "fmt"
)

type UserController struct {
    beego.Controller
}


func (self *UserController) Get() {
    users, _ := userMgr.GetUserList()
    tr := rpc.TableResult{}
    tr.Rows =users
    tr.Total = len(users)
    self.Data["json"] = &tr
    self.ServeJson();
}

func (self *UserController) Post() {
    fmt.Println(self.Ctx.Input.RequestBody)
    fmt.Println(self.Input())
    ret := rpc.JsonResult{"success", "id"}
    self.Data["json"] = &ret
    self.ServeJson();
}