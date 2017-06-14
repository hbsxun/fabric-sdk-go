package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/chaincode"
	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	"github.com/spf13/pflag"
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
	flags := &pflag.FlagSet{}
	flags.StringVar(&common.ChaincodeID, common.ChaincodeIDFlag, req.ChaincodeName, "The unique name of chaincode")
	flags.StringVar(&common.ChaincodePath, common.ChaincodePathFlag, common.CCPathPrefix+req.ChaincodeName, "The source code path of chaincode")
	action, err := chaincode.NewInstallAction(flags)
	if err != nil {
		fmt.Printf("InstallCC Initialize error...")
	}
	err = action.Invoke()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("Install chaincode [%s] successful\n", req.ChaincodeName)
	}

	u.ServeJSON()
}
