package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/chaincode"
)

// Operations about ChaincodeInfo
type ChaincodeInfoController struct {
	beego.Controller
}

// @Title ChaincodeInfo
// @Description ChaincodeInfo chaincode on peers
// @Param	body		body	chaincode.ChaincodeInfoArgs  true		"body for chaincode Description"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (u *ChaincodeInfoController) Post() {
	var req chaincode.ChaincodeInfoArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := chaincode.NewChaincodeInfoAction(&req)
	if err != nil {
		fmt.Printf("ChaincodeInfo Initialize error...")
	}
	resp, err := action.Execute()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = resp
	}

	u.ServeJSON()
}
