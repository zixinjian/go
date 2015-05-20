package userMgr
import (
    "github.com/astaxie/beego/orm"
//    "fmt"
//    "github.com/astaxie/beego/logs"
)

type User struct {
    Id string `orm:"pk;column(id);" json:"id" beedb:"PK" form:"_"`
    Username string `json:"username" form:"_"`
    Password string `json:"password" form:"_"`
    Name string     `json:"name" form:"_"`
    Department string `json:"department" form:"_"`
    Role string       `json:"role" form:"_"`
    Flag string       `json:"flag" form:"_"`
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
//    log := logs.NewLogger(10000)
//    log.SetLogger("console", "")
//    log.Error("getuser:", username, password)

    qs := o.QueryTable("user")
    var user User
    err := qs.Filter("username", username).Filter("password", password).One(&user, "username", "role")
    //fmt.Println("err", err)
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
func AddUser(user User)(string, string){
//    o := orm.NewOrm()
//    user := new User{}
    return "error", ""
}
