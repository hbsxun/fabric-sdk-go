package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/chaincode"
)

// Operations about Invoke
type InvokeController struct {
	beego.Controller
}

// @Title Invoke
// @Description Invoke chaincode on peers
// @Param	body		body	chaincode.InvokeArgs  true		"body for chaincode Description"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router / [post]
func (u *InvokeController) Post() {
	var req chaincode.InvokeArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := chaincode.NewInvokeAction(&req)
	if err != nil {
		fmt.Printf("Invoke Initialize error...")
	}
	resp, err := action.Execute()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = resp
	}

	u.ServeJSON()
}
