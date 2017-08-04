package query

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/query"
)

// Operations about QueryInfo
type QueryInfoController struct {
	beego.Controller
}

// @Title QueryInfo
// @Description Query Info
// @Param	body		body 	query.QueryInfoArgs		true	"body for Query BlockChain Info"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (u *QueryInfoController) Post() {
	var req query.QueryInfoArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := query.NewQueryInfoAction(&req)
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
