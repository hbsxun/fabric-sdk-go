package query

import (
	"fmt"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/query"
)

// Operations about QueryInstalled
type QueryTxController struct {
	beego.Controller
}

// @Title QueryTx
// @Description query transaction
// @Param	body		body	query.QueryTxArgs   true		"body for querytx Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router / [post]
func (u *QueryTxController) Post() {
	var req query.QueryTxArgs

	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)

	action, err := query.NewQueryTXAction(&req)
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
