package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"webo/controllers/ui"
)

type UiController struct {
	beego.Controller
}

func (this *UiController) Add() {
	fmt.Println("requestBosy", this.Ctx.Input.RequestBody)
	fmt.Println("params", this.Ctx.Input.Params)
	fmt.Println("requestBosy", this.Input()["id"])
	item, ok := this.Ctx.Input.Params[":hi"]
	if !ok {
		fmt.Println("hi", item)
	}

	this.Data["Service"] = item + "Service"
	this.Data["Form"]=ui.BuildAddForm(item)
	this.TplNames = "add.html"
}


func (this *UiController) Update(){
	fmt.Println("Update", this.Ctx.Input.RequestBody)
	fmt.Println("params", this.Ctx.Input.Params)
	fmt.Println("requestBosy", this.Input()["id"])
	item, ok := this.Ctx.Input.Params[":hi"]
	if !ok {
		fmt.Println("hi", item)
	}

	this.Data["Service"] = item + "Service"
	this.Data["Form"]=ui.BuildAddForm(item)
	this.TplNames = "add.html"
}