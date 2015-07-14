package controllers
import (
    "github.com/astaxie/beego"
    "webo/models/userMgr"
    "webo/models/rpc"
    "fmt"
)

type LoginController struct {
    beego.Controller
}

func (this *LoginController) Get() {
//    role := this.GetSession("role")
//    if role == nil {
    this.Redirect("/static/frame/login.html", 302)
//    } else {
//        this.Redirect("/static/main.html", 302)
//    }
}
func (this *LoginController) Post() {
    username := this.GetString("login_username")
    password := this.GetString("login_password")

    loginRet := rpc.JsonResult{}
    if username == "" || password == "" {
        loginRet.Result="请输入用户名和密码！"
    }

    user, err := userMgr.GetUser(username, password)
    if err != nil {
        fmt.Println("err", err)
        loginRet.Result="用户名或密码错误！"
    } else {
        this.SetSession("username", user.Username)
        this.SetSession("role", user.Role)
        loginRet.Ret="success"
    }
    this.Data["json"] = &loginRet
    this.ServeJson()
}

type LogoutController struct {
    beego.Controller
}

func (this *LogoutController) Get() {
//    fmt.Println("logout")
    this.DelSession("role")
    this.Redirect("/", 302)
}

