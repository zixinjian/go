package ui
import (
    "webo/models/entityDef"
    "fmt"
)

type FormBuilder struct {

}
var textTemp = `    <div class="form-group">
        <label class="col-sm-3 control-label">%s</label>
        <div class="col-sm-6">
            <input type="text" class="input-block-level form-control" data-validate="{required: true, messages:{required:'%s'}}" name="role" id="role" autocomplete="off" />
        </div>
    </div>`

func BuildForm(entity string)string{
    oEntityDef, ok := entityDef.EntityDefMap[entity]
    if ok{
        fmt.Println("abc", oEntityDef)
    }
    return
}