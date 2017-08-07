package ledger

import (
	"fmt"

	"encoding/json"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/query"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/ledger"
)

// @Title QueryBlocks
// @Description Query blocks
// @Param	body		body	query.QueryBlockArgs   true		"body for querytx Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /QueryBlocks [post]
func (u *LedgerController) QueryBlocks() {
	var req query.QueryBlockArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 307
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		blocksInfo, err := ledger.QueryBlocks(&req)
		if err != nil {
			res["status"] = 331
			res["message"] = fmt.Sprintf("QueryBlocks failed [%s]", err)
		} else {
			res["status"] = 200
			res["message"] = blocksInfo
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
