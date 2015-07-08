package svc

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"io/ioutil"
	"strings"
	"webo/models/util"
)

type field struct {
	Name    string      `json:name`
	Type    string      `json:type`
	Model   string      `json:model`
	Enum    []string    `json:enum`
	Default interface{} `json:default`
}

type entityDef struct {
	Name   string  `json:name`
	Fields []field `json:fields`
}

func (self *entityDef) GetFields() []string {
	fields := make([]string, len(self.Fields))
	for idx, attr := range self.Fields {
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
func List(entity string, params SvcParams) SvcResults {
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

	return SvcResults{"success", retList}
}

func Get(entity string, params SvcParams) (string, map[string]interface{}) {
	_, lstRet := List(entity, params)

	if len(lstRet) > 0 {
		return "success", lstRet[0]
	}
	return "success", nil
}
func Add(entity string, params SvcParams) string {

	Q := "'"
	oEntityDef, ok := entityDefMap[entity]
	if !ok {
		return "entity_no_define"
	}
	nFieldLen := len(oEntityDef.Fields)
	fields := make([]string, nFieldLen)
	marks := make([]string, nFieldLen)
	values := make([]interface{}, nFieldLen)
	for idx, attr := range oEntityDef.Fields {
		fields[idx] = attr.Name
		marks[idx] = "?"
		value, ok := params[attr.Name]
		if ok {
			values[idx] = value
			continue
		}
		if attr.Model == "key" {
			values[idx] = util.TUId()
			continue
		}
		values[idx] = attr.Default
	}

	fmt.Println(marks)

	sep := fmt.Sprintf("%s, %s", Q, Q)
	qmarks := strings.Join(marks, ", ")
	fmt.Println("qmarks:", qmarks)
	columns := strings.Join(fields, sep)
	fmt.Println("columns:", columns)

	query := fmt.Sprintf("INSERT INTO %s%s%s (%s%s%s) VALUES (%s)", Q, entity, Q, Q, columns, Q, qmarks)
	fmt.Println("query:", query)

	fmt.Println("values", values)
	o := orm.NewOrm()
	if res, err := o.Raw(query, values...).Exec(); err == nil {
		fmt.Println("res", res)
	}
	return "success"
}

func update(entity string, params SvcParams) string {
	Q := "'"
	oEntityDef, ok := entityDefMap[entity]
	if !ok {
		return "entity_no_define"
	}

	id, ok := params["id"]
	if !ok {
		return "no_id"
	}
	var names []string
	var values []interface{}
	for _, field := range oEntityDef.Fields {
		if field.Name == "id" {
			continue
		}
		value, ok := params[field.Name]
		if ok {
			values = append(values, value)
			names = append(names, field.Name)
		}
	}
	values = append(values, id)

	sep := fmt.Sprintf("%s = ?, %s", Q, Q)
	setColumns := strings.Join(names, sep)

	query := fmt.Sprintf("UPDATE %s%s%s SET %s%s%s = ? WHERE %s%s%s = ?", Q, entity, Q, Q, setColumns, Q, Q, "id", Q)

	o := orm.NewOrm()
	if res, err := o.Raw(query, values...).Exec(); err == nil {
		fmt.Println("res", res)
	}
	return "success"
}
