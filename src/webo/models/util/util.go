package util
import (
    "time"
    "fmt"
)

var gId uint32
var gOldTime time.Time



func TUId() string {
    now:=time.Now()
    if(gOldTime.After(now)) {
        now = gOldTime
    }
    if(gId > 999){
        gId = 0
        now.Add(time.Second)
    }
    gOldTime = now
//    fmt.Println(now.Format("20060102150405"))
    ret:=fmt.Sprintf("%s%03d", now.Format("20060102150405"), gId)
    gId = gId + 1
    return ret
}


func init() {
    gId=0
    gOldTime=time.Now()
}

