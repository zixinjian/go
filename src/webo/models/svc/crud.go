package svc

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"strings"
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
func List(name string) []map[string]interface{}{
	o := orm.NewOrm()
	var maps []orm.Params
	o.QueryTable(name).Values(&maps)
	retList:=make([]map[string]interface{}, len(maps))

	for idx, oldMap :=range maps {
		var retMap = make(map[string]interface{}, len(oldMap))
		for key, value:= range oldMap {
			retMap[strings.ToLower(key)]=value
		}
		retList[idx] = retMap
	}
	return retList
//	fmt.Println(len(retList), retList)
}

//func serve(name string, method string) interface{}{
//
//}
