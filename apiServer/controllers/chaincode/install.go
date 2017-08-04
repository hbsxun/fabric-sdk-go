package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/chaincode"
)

// Operations about InstallCC
type InstallCCController struct {
	beego.Controller
}

// @Title InstallCC
// @Description InstallCC on peers
// @Param	body		body	chaincode.InstallCCArgs   true		"body for chaincode Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router / [post]
func (u *InstallCCController) Post() {
	var req chaincode.InstallCCArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := chaincode.NewInstallAction(&req)
	if err != nil {
		fmt.Printf("InstallCC Initialize error...")
	}
	err = action.Execute()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("Install chaincode [%s] successful\n", req.ChaincodeName)
	}

	u.ServeJSON()
}
