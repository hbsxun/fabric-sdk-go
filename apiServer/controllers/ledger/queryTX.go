package ledger

import (
	"fmt"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/ledger"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/query"
)

// Operations about Ledger
type LedgerController struct {
	beego.Controller
}

// @Title QueryTX
// @Description Query transaction
// @Param	body		body	query.QueryTxArgs   true		"body for querytx Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /QueryTX [post]
func (u *LedgerController) QueryTX() {
	var req query.QueryTxArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 307
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		txInfo, err := ledger.QueryTX(&req)
		if err != nil {
			res["status"] = 330
			res["message"] = fmt.Sprintf("QueryTX failed [%s]", err)
		} else {
			res["status"] = 200
			res["message"] = txInfo
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
