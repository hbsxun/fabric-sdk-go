package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/chaincode"
)

// Operations about Instantiate
type InstantiateController struct {
	beego.Controller
}

// @Title Instantiate
// @Description Instantiate chaincode on peers
// @Param	body		body	chaincode.InstantiateArgs  true		"body for chaincode Description"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router / [post]
func (u *InstantiateController) Post() {
	var req chaincode.InstantiateArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := chaincode.NewInstantiateAction(&req)
	if err != nil {
		fmt.Printf("Instantiate Initialize error...")
	}
	resp, err := action.Execute()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = resp
	}

	u.ServeJSON()
}
