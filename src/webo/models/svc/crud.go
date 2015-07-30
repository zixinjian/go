package svc

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"webo/models/itemDef"
	"webo/models/util"
)

type SelectBuilder struct {
	Table		string

}


func List(entity string, queryParams Params, limitParams map[string]int64, orderBy Params) (string, int64, []map[string]interface{}) {
	oEntityDef, ok := itemDef.EntityDefMap[entity]
	retList := make([]map[string]interface{}, 0)
	if !ok {
		return "entity_no_define", 0, retList
	}
	o := orm.NewOrm()



	qs := o.QueryTable(entity)
	for key, value := range queryParams {
		qs = qs.Filter(key, value)
	}
	total, err:=qs.Count()
	if err != nil{
		return "count_total_error", 0, retList
	}
	//fmt.Println("total", total)
	if limit, ok:=limitParams["limit"]; ok{
		//fmt.Println("limit", limit)
		qs = qs.Limit(limit)
	}
	if offset, ok:=limitParams["offset"]; ok{
		//fmt.Println("offset", offset)
		qs = qs.Offset(offset)
	}

	var orderList []string
	fieldMap := oEntityDef.GetFieldMap()
	for k, v:= range orderBy{
		_, ok := fieldMap[k]
		if !ok {
			continue
		}

		if strings.EqualFold(v.(string), "asc"){
			orderList = append(orderList, k)
		}
		if strings.EqualFold(v.(string), "desc"){
			orderList = append(orderList, "-" + v.(string))
		}
	}
	//fmt.Println("orderBy", orderList)
	if len(orderList) > 0{
		qs = qs.OrderBy(orderList...)
	}
	var resultMaps []orm.Params
	qs.Values(&resultMaps)

	retList = make([]map[string]interface{}, len(resultMaps))
	fmt.Println("old", resultMaps)
	for idx, oldMap := range resultMaps {
		var retMap = make(map[string]interface{}, len(oldMap))
		for key, value := range oldMap {
			retMap[strings.ToLower(key)] = value
		}
		retList[idx] = retMap
	}
	return "success", total, retList
}

func Get(entity string, params Params) (string, map[string]interface{}) {
	_, _, retList := List(entity, params, map[string]int64{}, Params{})
	if len(retList) > 0 {
		return "success", retList[0]
	}
	return "not_found", nil
}

func Add(entity string, params Params) string {

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
		if attr.Model == "sn" {
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
//
//	fmt.Println("values", values)
	o := orm.NewOrm()
	if res, err := o.Raw(query, values...).Exec(); err != nil {
		fmt.Println(err)
		fmt.Println("res", res)
	}
	return "success"
}

func Update(entity string, params Params) string {
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
