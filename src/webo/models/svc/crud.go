package svc

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"strings"
	//	"webo/models/util"
)

type attribute struct {
	Name string   `json:name`
	Type string   `json:type`
	Enum []string `json:enum`
}

type entityDef struct {
	Name       string      `json:name`
	Attributes []attribute `json:attributes`
}

func (self *entityDef) GetFields() []string {
	fields := make([]string, len(self.Attributes))
	for idx, attr := range self.Attributes {
		fields[idx] = attr.Name
	}
	return fields
}

var entityDefMap = make(map[string]entityDef)

func init() {
	bytes, err := ioutil.ReadFile("conf/entity.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
	}
	//    fmt.Println("readFile", bytes)
	if err := json.Unmarshal(bytes, &entityDefMap); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
	}
	fmt.Println(entityDefMap)
	//    list("user")
}
func List(entity string, params Params) (string, []map[string]interface{}) {
	o := orm.NewOrm()
	var resultMaps []orm.Params
	qs := o.QueryTable(entity)
	for key, value := range params {
		qs = qs.Filter(key, value)
	}
	qs.Values(&resultMaps)

	retList := make([]map[string]interface{}, len(resultMaps))
	for idx, oldMap := range resultMaps {
		var retMap = make(map[string]interface{}, len(oldMap))
		for key, value := range oldMap {
			retMap[strings.ToLower(key)] = value
		}
		retList[idx] = retMap
	}
	return "success", retList
}

func Get(entity string, params Params) (string, map[string]interface{}) {
	_, lstRet := List(entity, params)

	if len(lstRet) > 0 {
		return "success", lstRet[0]
	}
	return "success", nil
}
func Add(entity string, params Params) string {
	//	o:=orm.NewOrm()
	Q := "'"
	oEntityDef, ok := entityDefMap[entity]
	if !ok {
		return "entity_no_define"
	}
	fields := oEntityDef.GetFields()
	marks := make([]string, len(fields))
	for i, _ := range marks {
		marks[i] = "?"
	}
	fmt.Println(marks)

	sep := fmt.Sprintf("%s, %s", Q, Q)
	qmarks := strings.Join(marks, ", ")
	fmt.Println("qmarks:", qmarks)
	columns := strings.Join(fields, sep)
	fmt.Println("columns:", columns)

	query := fmt.Sprintf("INSERT INTO %s%s%s (%s%s%s) VALUES (%s)", Q, entity, Q, Q, columns, Q, qmarks)
	fmt.Println("query:", query)
	//	res, err := o.Raw("INSERT INTO %s%s%s (%s%s%s) VALUES (%s)", util.TUId(), "your").Exec()
	//	fmt.Println(res, err)
	return "success"
}

func CreateID() {

}
