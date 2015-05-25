package svc
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)


type attribute struct {
    Name string `json:name`
    Type string `json:type`
    Enum []string `json:enum`
}


type entity struct {
    Name string `json:name`
    Attributes map[string]attribute `json:attributes`
}

type entityMap map[string]entity

func getEntity(entity string) ([]byte, error){
    entity, ok := entityMap[entity]
    if ok{
        return entity
    }
    return nil;
}

func init() {
    bytes, err := ioutil.ReadFile("conf/entity.json")
    if err != nil {
        fmt.Println("ReadFile: ", err.Error())
        return nil, err
    }
    fmt.Println("readFile", bytes)
    if err := json.Unmarshal(bytes, &entityMap); err != nil {
        fmt.Println("Unmarshal: ", err.Error())
        return nil, err
    }
    return nil, nil
}