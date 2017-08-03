package query

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/query"
)

// Operations about QueryBlock
type QueryController struct {
	beego.Controller
}

// @Title QueryBlock
// @Description Query Block
// @Param	body		body 	query.QueryBlockArgs		true	"body for Query Block"
// @Success 200 {body}
// @Failure 403 body is empty
// @router /QueryBlock [post]
func (u *QueryController) Post() {
	var req query.QueryBlockArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 80401
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := query.NewQueryBlockAction(&req)
		if err != nil {
			fmt.Printf("QueryBlock Initialize error...")
			res["status"] = 80402
			res["message"] = fmt.Sprintf("QueryBlock action error [%s]", err)
		} else {
			resp, err := action.Execute()
			if err != nil {
				res["status"] = 80403
				res["message"] = fmt.Sprintf("QueryBlock execute error [%s]", err)
			} else {
				res["status"] = 80200
				res["message"] = resp
			}
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
