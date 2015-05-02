package userMgr
import (
    "github.com/astaxie/beego/orm"
    "fmt"
    "github.com/astaxie/beego/logs"
)

type User struct {
    Id string `orm:"pk;column(id);" json:"id" beedb:"PK"`
    Username string `json:"username"`
    Password string `json:"password"`
    Name string     `json:"name"`
    Department string `json:"department"`
    Role string       `json:"role"`
    Flag string       `json:"flag"`
}

//func (this *User) TableName() string {
//    return "user"
//}

func init() {
    orm.RegisterModel(new(User))
}

func GetUser(username string, password string) (User, error) {
    orm.Debug = true
    o := orm.NewOrm()
    log := logs.NewLogger(10000)
    log.SetLogger("console", "")
    log.Error("getuser:", username, password)

    qs := o.QueryTable("user")
    var user User
    err := qs.Filter("username", username).Filter("password", password).One(&user, "username", "role")
    fmt.Println("err", err)
    if(err == nil){
        return user, nil
    }
    return user, err
}


func GetUserList()([]*User, string){
    o := orm.NewOrm()
    o.Using("default")
    var users []*User
    _, err := o.QueryTable("user").All(&users)
    if err == orm.ErrNoRows {
        return users, "success"
    }
    if err !=nil{
        return users, "error"
    }
    for _, user := range(users){
        user.Password = ""
    }
    return users, "success"
}
