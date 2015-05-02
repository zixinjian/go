package main

import (
	_ "webo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/astaxie/beego/context"
	"strings"
	"fmt"
)
func initDb(){
	orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	orm.RegisterDataBase("default", "sqlite3", "db/frame.sqlite")
}

var FilterUser = func(ctx *context.Context) {
	fmt.Println("url", ctx.Request.RequestURI)
	if strings.HasPrefix(ctx.Request.RequestURI, "/asserts"){
		fmt.Println("assert")
		return
	}
	if ctx.Request.RequestURI == "/static/frame/login.html"{
		fmt.Println("login")
		return
	}
	_, ok := ctx.Input.Session("role").(string)
	if !ok && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302, "/login")
	}
}


func main() {
	initDb()
	beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
	beego.Run()
}

