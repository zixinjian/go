package controllers

import (
    "github.com/astaxie/beego"
    "encoding/json"
    "fmt"
)

type ServiceController struct {
    beego.Controller
}
type jsonRespons struct {
    Data []string `json:"data"`
    Total int64 `json:"total"`
}

type jsonRequest struct {
    Id     *json.RawMessage `json:"id"`
    Method string           `json:"method"`
    Params *json.RawMessage `json:"params"`
}

type jsonResponse struct {
    Id     *json.RawMessage `json:"id"`
    Result interface{}      `json:"result"`
    Error  interface{}      `json:"error"`
}

type tableResult struct {
    Total int32 `json:"total"`
    Rows []map[string]string `json:"rows"`
}

func (this *ServiceController) Get() {
    fmt.Println("requestBosy", this.Ctx.Input.RequestBody)
    resp := new(jsonResponse)
    tr := new(tableResult)
    tr.Rows = []map[string]string{{"id":"1", "user":"user1", "name":"a", "department":"dep1", "role":"admin", "flat":""}}
    tr.Total = 1
    resp.Result = tr
    resp.Error = nil
    this.Data["json"] = tr
    this.ServeJson()
}
func (this *ServiceController) Post() {
    json.Unmarshal(this.Ctx.Input.RequestBody, &map[string]string{})
    fmt.Println("requestBosy", this.Ctx.Input.RequestBody)
    this.Data["json"] = "{\"ObjectId\":\"abcdID\"}"
    this.ServeJson();
}