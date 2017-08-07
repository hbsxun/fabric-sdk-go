package ledger

import (
	"fmt"

	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/query"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/ledger"
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
		res["status"] = 80401
		res["message"] = err.Error()
	} else {
		fmt.Println(req)
		txInfo, err := ledger.QueryTX(&req)
		if err != nil {
			res["status"] = 80402
			res["message"] = err.Error()
		} else {
			res["status"] = 80200
			res["message"] = txInfo
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}
