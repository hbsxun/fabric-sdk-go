package ledger

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/query"
)

// @Title QueryInfo
// @Description Query information of peerUrl in the channel
// @Param	body		body 	query.QueryInfoArgs		true	"body for Query BlockChain Info"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /QueryInfo [post]
func (u *LedgerController) QueryInfo() {
	var req query.QueryInfoArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 307
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err.Error())
	} else {
		fmt.Println(req)
		action, err := query.NewQueryInfoAction(&req)
		if err != nil {
			res["status"] = 334
			res["message"] = fmt.Sprintf("NewQueryInfoAction failed [%s]", err.Error())
		} else {
			resp, err := action.Execute()
			if err != nil {
				res["status"] = 334
				res["message"] = fmt.Sprintf("QueryInfo execute failed [%s]", err.Error())
			} else {
				res["status"] = 200
				res["message"] = resp
			}
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
