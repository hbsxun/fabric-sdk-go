package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/fabric/chaincode"
)

// Operations about Query
type QueryController struct {
	beego.Controller
}

// @Title Query
// @Description Query chaincode on peers
// @Param	body		body	chaincode.QueryArgs  true		"body for chaincode Description"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router / [post]
func (u *QueryController) Post() {
	var req chaincode.QueryArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := chaincode.NewQueryAction(&req)
	if err != nil {
		fmt.Printf("Query Initialize error...")
	}
	resp, err := action.Execute()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = resp
	}

	u.ServeJSON()
}
