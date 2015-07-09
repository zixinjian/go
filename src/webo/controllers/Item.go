package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"webo/models/svc"
)

type TableResult struct {
	Status string      `json:status`
	Total  int64       `json:"total"`
	Rows   interface{} `json:"rows"`
}

type ItemController struct {
	beego.Controller
}

func (this *ItemController) List() {
	fmt.Println("requestBosy", this.Ctx.Input.RequestBody)
	fmt.Println("params", this.Ctx.Input.Params)
	fmt.Println("requestBosy", this.Input()["id"])
	fmt.Println("ge", this.GetString("xx"))
	hi, ok := this.Ctx.Input.Params[":hi"]
	if !ok {
		this.Data["json"] = TableResult{"success", 0, ""}
	}
	params := make(svc.SvcParams)
	params["username"] = "a"
	result, total, retList := svc.List(hi, params)
	fmt.Println(result, total, retList)
	this.Data["json"] = &TableResult{result, int64(total), retList}
	this.ServeJson()
}
func (this *ItemController) Get() {
	fmt.Println("requestBosy", this.Ctx.Input.RequestBody)
	fmt.Println("params", this.Ctx.Input.Params)
	tr := new(TableResult)
	tr.Rows = []map[string]string{{"id": "1", "user": "user1", "name": "a", "department": "dep1", "role": "admin", "flat": ""}}
	tr.Total = 1
	this.Data["json"] = tr
	this.ServeJson()
}
func (this *ItemController) Add() {
	fmt.Println("requestBosy", this.Ctx.Input.RequestBody)
	fmt.Println("params", this.Ctx.Input.Params)
	tr := new(TableResult)
	tr.Rows = []map[string]string{{"id": "1", "user": "user1", "name": "a", "department": "dep1", "role": "admin", "flat": ""}}
	tr.Total = 1
	this.Data["json"] = tr
	this.ServeJson()
}
func (this *ItemController) Update() {
	fmt.Println("requestBosy", this.Ctx.Input.RequestBody)
	fmt.Println("params", this.Ctx.Input.Params)
	tr := new(TableResult)
	tr.Rows = []map[string]string{{"id": "1", "user": "user1", "name": "a", "department": "dep1", "role": "admin", "flat": ""}}
	tr.Total = 1
	this.Data["json"] = tr
	this.ServeJson()
}
func (this *ItemController) Delete() {
	fmt.Println("requestBosy", this.Ctx.Input.RequestBody)
	fmt.Println("params", this.Ctx.Input.Params)
	tr := new(TableResult)
	tr.Rows = []map[string]string{{"id": "1", "user": "user1", "name": "a", "department": "dep1", "role": "admin", "flat": ""}}
	tr.Total = 1
	this.Data["json"] = tr
	this.ServeJson()
}
