package itemDef

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Field struct {
	Name    string      `json:name`
	Type    string      `json:type`
	Label   string      `json:label`
	Input   string      `json:input`
	Require string      `json:require`
	Model   string      `json:model`
	Enum    []string    `json:enum`
	Default interface{} `json:default`
}

type ItemDef struct {
	Name   string  `json:name`
	Fields []Field `json:fields`
}

func (self *ItemDef) GetFields() []string {
	fields := make([]string, len(self.Fields))
	for idx, field := range self.Fields {
		fields[idx] = field.Name
	}
	return fields
}

var EntityDefMap = make(map[string]ItemDef)

func init() {
	fmt.Println("initItemDefMap")
	bytes, err := ioutil.ReadFile("conf/item.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
	}
	//    fmt.Println("readFile", bytes)
	if err := json.Unmarshal(bytes, &EntityDefMap); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
	}
	fmt.Println(EntityDefMap)
	//    list("user")
}
