package controllers

import (
    "github.com/astaxie/beego"
    "webo/models/userMgr"
    "webo/models"
)

type UserController struct {
    beego.Controller
}


func (self *UserController) Get() {
//    users := [...]map[string]string{{"id":"1", "user":"user1", "name":"a", "department":"dep1", "role":"admin", "flat":""}}
    users, _ := userMgr.GetUserList()
    tr := models.TableResult{}
    tr.Rows =users
    tr.Total = len(users)
    self.Data["json"] = &tr
    self.ServeJson();
}

func (self *UserController) Post() {
    //user := user{"1", "user1", "a", "name1", "dep1", "admin", ""}
    users := map[string]string{"id":"1", "user":"user1", "name":"a", "department":"dep1", "role":"admin", "flat":""}
    self.Data["json"] = &users

    self.ServeJson();
}