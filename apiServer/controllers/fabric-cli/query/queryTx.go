package query

import (
	"fmt"

	"encoding/json"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/query"
)

// @Title QueryTx
// @Description query transaction
// @Param	body		body	query.QueryTxArgs   true		"body for querytx Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /QueryTx [post]
func (u *QueryController) QueryTx() {
	var req query.QueryTxArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 301
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := query.NewQueryTXAction(&req)
		if err != nil {
			fmt.Printf("QueryTx Initialize error...")
			res["status"] = 404
			res["message"] = fmt.Sprintf("QueryTx action error [%s]", err)
		} else {
			err := action.Execute()
			if err != nil {
				res["status"] = 404
				res["message"] = fmt.Sprintf("QueryTx execute error [%s]", err)
			} else {
				res["status"] = 200
				res["message"] = fmt.Sprintf("query tx [%s] successfully", req.TxID)
			}
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
