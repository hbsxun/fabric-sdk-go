package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/chaincode"
)

// Operations about ChaincodeInfo
type ChaincodeController struct {
	beego.Controller
}

// @Title ChaincodeInfo
// @Description ChaincodeInfo chaincode on peers
// @Param	body		body	chaincode.ChaincodeInfoArgs  true		"body for chaincode Description"
// @Success 200 {body}
// @Failure 403 body is empty
// @router /ChaincodeInfo [post]
func (u *ChaincodeController) Post() {
	var req chaincode.ChaincodeInfoArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 80401
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := chaincode.NewChaincodeInfoAction(&req)
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("NewChaincodeInfoAction failed [%s]", err)
		} else {
			resp, err := action.Execute()
			if err != nil {
				res["status"] = 80403
				res["message"] = fmt.Sprintf("GetChaincodeInfo Execute error [%s]", err)
			} else {
				res["status"] = 80200
				res["message"] = resp
			}
		}
	}

	u.Data["json"] = res
	u.ServeJSON()
}
