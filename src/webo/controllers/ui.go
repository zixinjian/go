package controllers

import (
    "github.com/astaxie/beego"
    "fmt"
)

type UiController struct {
    beego.Controller
}
func (this *UiController) Add(){
    fmt.Println("requestBosy", this.Ctx.Input.RequestBody)
    fmt.Println("params", this.Ctx.Input.Params)
    fmt.Println("requestBosy", this.Input()["id"])
    fmt.Println("ge", this.GetString("xx"))
    hi, ok := this.Ctx.Input.Params[":hi"]
    if !ok {
        fmt.Println("hi", hi)
    }
    this.Data["Service"] = "userService"
    this.Data["Method"] = "add"
    this.TplNames = "add.html"
}
