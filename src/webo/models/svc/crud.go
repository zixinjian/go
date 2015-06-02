package svc

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
)

type attribute struct {
	Name string   `json:name`
	Type string   `json:type`
	Enum []string `json:enum`
}

type entity struct {
	Name       string      `json:name`
	Attributes []attribute `json:attributes`
}

var entityMap = make(map[string]entity)

func init() {
	bytes, err := ioutil.ReadFile("conf/entity.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
	}
	//    fmt.Println("readFile", bytes)
	if err := json.Unmarshal(bytes, &entityMap); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
	}
	//    fmt.Println(entityMap)
	//    list("user")
}
func List(name string) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.QueryTable("user").Values(&maps)
	fmt.Println(maps)
}

//func serve(name string, method string) interface{}{
//
//}
