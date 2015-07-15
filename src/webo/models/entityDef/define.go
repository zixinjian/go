package entityDef
import(
    "io/ioutil"
    "fmt"
    "encoding/json"
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

var EntityDefMap = make(map[string]entityDef)

func init() {
    fmt.Println("initEntityMap")
    bytes, err := ioutil.ReadFile("conf/entity.json")
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