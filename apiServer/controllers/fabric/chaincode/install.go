package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/chaincode"
)

// @Title InstallCC
// @Description InstallCC on peers
// @Param	body		body	chaincode.InstallCCArgs   true		"body for chaincode Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /InstallCC [post]
func (u *ChaincodeController) InstallCC() {
	var req chaincode.InstallCCArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 80401
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := chaincode.NewInstallAction(&req)
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("NewInstallAction failed[%s]", err)
		} else {
			err = action.Execute()
			if err != nil {
				res["status"] = 80403
				res["message"] = fmt.Sprintf("InstallCC execute error [%s]", err)
			} else {
				res["status"] = 80200
				res["message"] = fmt.Sprintf("Install chaincode [%s] successfully", req.ChaincodeID)
			}
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}
