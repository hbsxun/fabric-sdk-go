package query

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/query"
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
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 301
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := query.NewQueryChainInfoAction(&req)
		if err != nil {
			fmt.Printf("QueryChainInfo Initialize error...")
			res["status"] = 402
			res["message"] = fmt.Sprintf("QueryChainInfo action error [%s]", err)
		} else {
			err := action.Execute()
			if err != nil {
				res["status"] = 402
				res["message"] = fmt.Sprintf("QueryChainInfo execute error [%s]", err)
			} else {
				res["status"] = 200
				res["message"] = "query chain info successfully"
			}
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
