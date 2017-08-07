package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/chaincode"
)

// @Title QueryCC
// @Description Query chaincode on peers
// @Param	body		body	chaincode.QueryArgs  true		"body for chaincode Description"
// @Success 200 {body}
// @Failure 403 body is empty
// @router /QueryCC [post]
func (u *ChaincodeController) QueryCC() {
	var req chaincode.QueryArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 80401
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := chaincode.NewQueryAction(&req)
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("NewQueryAction failed[%s]", err)
		} else {
			resp, err := action.Query()
			if err != nil {
				res["status"] = 80403
				res["message"] = fmt.Sprintf("QueryCC execute error [%s]", err)
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
