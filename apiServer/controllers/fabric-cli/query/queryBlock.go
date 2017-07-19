package query

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/query"
)

// Operations about QueryBlock
type QueryBlockController struct {
	beego.Controller
}

// @Title QueryBlock
// @Description Query Block
// @Param	body		body 	query.QueryBlockArgs		true	"body for Query Block"
// @Success 200 {string} []string
// @Failure 403 body is empty
// @router / [post]
func (u *QueryBlockController) Post() {
	var req query.QueryBlockArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := query.NewQueryBlockAction(&req)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		resp, err := action.Execute()
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = resp
		}
	}

	u.ServeJSON()
}
