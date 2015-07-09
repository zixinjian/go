package controllers

import (
	"github.com/astaxie/beego"
	//    "encoding/json"
	"fmt"
	"webo/models/svc"
	"webo/models/rpc"
)

type ServiceController struct {
	beego.Controller
}

func (this *ServiceController) Get() {
	//    fmt.Println("requestBosy", this.Ctx.Input.RequestBody)
	//    resp := new(jsonResponse)
	//    tr := new(tableResult)
	//    tr.Rows = []map[string]string{{"id":"1", "user":"user1", "name":"a", "department":"dep1", "role":"admin", "flat":""}}
	//    tr.Total = 1
	//    resp.Result = tr
	//    resp.Error = nil
	//    this.Data["json"] = tr
	//    this.ServeJson()
}

func (this *ServiceController) Post() {
	//    json.Unmarshal(this.Ctx.Input.RequestBody, &map[string]string{})
	fmt.Println("requestBosy", this.Input())
	inputMap := this.Input()
	crud, ok := inputMap["crud"]
	if ok {
		fmt.Println("crud", crud)
	}
	service, ok := inputMap["service"]
	if !ok {
		this.Data["json"] = &svc.SvcResults{"no_service_param", ""}
	}
	fmt.Println("service", service)
	method, ok := inputMap["method"]
	if !ok {
		this.Data["json"] = &svc.SvcResults{"no_method_param", ""}
	}
	fmt.Println("method", method)
	this.Data["json"] = "{\"ObjectId\":\"abcdID\"}"
	this.ServeJson()
}

func list(entity string, params svc.SvcParams) svc.SvcResults {
	result, total, retList := svc.List(entity, params)
	return svc.SvcResults{result, rpc.TableResult{total, retList}}
}