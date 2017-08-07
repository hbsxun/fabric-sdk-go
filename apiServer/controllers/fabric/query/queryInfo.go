package query

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/query"
)

// @Title QueryChainInfo
// @Description Query Chain Info
// @Param	body		body 	query.QueryChainInfoArgs		true	"body for Query BlockChain Info"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /QueryChainInfo [post]
func (u *QueryController) QueryChainInfo() {
	var req query.QueryChainInfoArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 80401
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := query.NewQueryChainInfoAction(&req)
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("NewQueryChainInfoAction failed[%s]", err)
		} else {
			resp, err := action.Execute()
			if err != nil {
				res["status"] = 80403
				res["message"] = fmt.Sprintf("QueryChainInfo execute error [%s]", err)
			} else {
				res["status"] = 80200
				res["message"] = resp
			}
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}
