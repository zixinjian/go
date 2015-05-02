package controllers
import (
    "github.com/astaxie/beego"
//    "webo/models/userMgr"
    "github.com/astaxie/beego/logs"
    "webo/models/userMgr"
)

type LoginController struct {
    beego.Controller
}

func (this *LoginController) Get() {
    role := this.GetSession("role")
    if role == nil {
        this.Redirect("/static/frame/login.html", 302)
    } else {
        this.Redirect("/static/main.html", 302)
    }
}
type loginResult struct {
    Ret string `json:"ret"`
    Result interface{} `json:"result"`
}
func (this *LoginController) Post() {
    username := this.GetString("login_username")
    password := this.GetString("login_password")
    log := logs.NewLogger(10000)
    log.SetLogger("console", "")
    log.Error("post:", username, password)

    loginRet := loginResult{}
    if username == "" || password == "" {
        loginRet.Result="请输入用户名和密码！"
    }

    user, err := userMgr.GetUser(username, password)
    if err != nil {
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
    this.DelSession("username")
    this.Redirect("/", 302)
}

