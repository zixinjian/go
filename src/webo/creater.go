package main

import (
	"fmt"
	"webo/models/itemDef"
)

func main() {
	itemDefMap := itemDef.EntityDefMap
	for item, def := range itemDefMap {
		fmt.Println(item, def)
	}
	fmt.Println("start")
}
