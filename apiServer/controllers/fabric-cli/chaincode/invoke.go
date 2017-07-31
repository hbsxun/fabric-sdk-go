package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/chaincode"
)

// @Title InvokeCC
// @Description Invoke chaincode on peers
// @Param	body		body	chaincode.InvokeArgs  true		"body for chaincode Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /InvokeCC [post]
func (u *ChaincodeController) InvokeCC() {
	var req chaincode.InvokeArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 301
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := chaincode.NewInvokeAction(&req)
		if err != nil {
			fmt.Printf("Invoke Initialize error...")
			res["status"] = 306
			res["message"] = fmt.Sprintf("Invoke action error [%s]", err)
		} else {
			err := action.Execute()
			if err != nil {
				res["status"] = 306
				res["message"] = fmt.Sprintf("Invoke execute error [%s]", err)
			} else {
				res["status"] = 200
				res["message"] = fmt.Sprintf("Invoke chaincode [%s] successfully", req.ChaincodeID)
			}
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
