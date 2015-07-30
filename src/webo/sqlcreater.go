package main
import (
    "webo/models/itemDef"
    "fmt"
)
func main() {
    for itemName, oItemDef := range itemDef.EntityDefMap{
//        fmt.Println(itemName, oItemDef)
        fieldsql := "CREATE TABLE " + itemName + " (id integer NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE"
        for _, field := range oItemDef.Fields{
//            fmt.Println("idx:", idx)
            fieldsql = fieldsql + "," + field.Name
            switch field.Model{
                case "sn", "text", "password", "enum":
                    fieldsql = fieldsql + " varchar"
                case "timestamp":
                    fieldsql = fieldsql + " timestamp"
                case "integer":
                    fieldsql = fieldsql + " integer"
                default:
                    fmt.Println("no such modal", field.Name, field.Model)
            }
            if field.Require == "true"{
                fieldsql = fieldsql + " NOT NULL"
            }
            if field.Unique == "true"{
                fieldsql = fieldsql + " UNIQUE"
            }
            if field.Default == ""{
                continue
            }
            switch field.Model{
                case "sn", "text", "password", "enum", "integer":
                fieldsql = fieldsql + " DEFAULT " + field.Default.(string)
                default:
                fmt.Println("no default", field.Name)
            }
        }

        fieldsql = fieldsql + ")"
        fmt.Println(fieldsql)
    }
}