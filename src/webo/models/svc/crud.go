package svc

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"webo/models/itemDef"
	"webo/models/util"
)

func List(entity string, params SvcParams) (string, int, []map[string]interface{}) {
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
	return "success", len(retList), retList
}

func Get(entity string, params SvcParams) (string, interface{}) {
	_, _, retList := List(entity, params)

	if len(retList) > 0 {
		return "success", retList[0]
	}
	return "success", nil
}

func Add(entity string, params SvcParams) string {

	Q := "'"
	oEntityDef, ok := itemDef.EntityDefMap[entity]
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

func Update(entity string, params SvcParams) string {
	Q := "'"
	oEntityDef, ok := itemDef.EntityDefMap[entity]
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

		if value, ok := params[field.Name]; ok {
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
