package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/chaincode"
)

// @Title InstantiateCC
// @Description Instantiate chaincode on peers
// @Param	body		body	chaincode.InstantiateArgs  true		"body for chaincode Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /InstantiateCC [post]
func (u *ChaincodeController) InstantiateCC() {
	var req chaincode.InstantiateArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 80401
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := chaincode.NewInstantiateAction(&req)
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("NewInstantiateAction failed[%s]", err)
		} else {
			err := action.Execute()
			if err != nil {
				res["status"] = 80403
				res["message"] = fmt.Sprintf("InstantiateCC Execute error [%s]", err)
			} else {
				res["status"] = 80200
				res["message"] = fmt.Sprintf("Instantiate chaincode [%s] successfully", req.ChaincodeID)
			}
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}
