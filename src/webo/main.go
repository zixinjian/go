package main

import (
//	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"webo/models/svc"
	_ "webo/routers"
	"fmt"
	"webo/models/util"
)

func initDb() {
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "db/frame.sqlite")
}

var FilterUser = func(ctx *context.Context) {
	//	fmt.Println("filterUser")
	//	fmt.Println("url", ctx.Request.RequestURI)
	//	if strings.HasPrefix(ctx.Request.RequestURI, "/asserts"){
	//		fmt.Println("assert")
	//		return
	//	}
	if ctx.Request.RequestURI == "/logout" {
		return
	}
	_, ok := ctx.Input.Session("role").(string)
	//	fmt.Println("role", role)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}

//var FilterStatic = func(ctx *context.Context){
////	fmt.Println("url", ctx.Request.RequestURI)
////	fmt.Println("FilterStatic")
//	if strings.HasPrefix(ctx.Request.RequestURI, "/asserts"){
////		fmt.Println("assert")
//		return
//	}
//	if ctx.Request.RequestURI == "/static/frame/login.html"{
////		fmt.Println("login")
//		return
//	}
//	role, ok := ctx.Input.Session("role").(string)
//	fmt.Println("role", role)
//	if !ok && ctx.Request.RequestURI != "/login" {
//		ctx.Redirect(302, "/login")
//	}
//}

func main() {
	initDb()
	//	beego.InsertFilter("/*", beego.BeforeStatic, FilterStatic)
//	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
//	fmt.Println(time.Now().Format("200601021504055"))
//	params := make(svc.Params)
//	params["password"]="a"
//	params["username"]="a"
	params := svc.Params{
		"username":"a",
		"password":"a",
	}
	fmt.Println(svc.List("user", params))
	fmt.Println(svc.Get("user", params))
	fmt.Println(util.TUId())
	fmt.Println(svc.Add("user", params))
//	beego.Run()
}
