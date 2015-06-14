package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"webo/models/svc"
	_ "webo/routers"
	//	"fmt"
	"fmt"
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
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	params := make(svc.Params)
	params["password"]="a"
	params["username"]="a"
	fmt.Println(svc.List("user", params))
	beego.Run()
}
