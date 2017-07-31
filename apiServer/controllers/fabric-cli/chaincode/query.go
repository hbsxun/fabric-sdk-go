package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/chaincode"
)

// @Title QueryCC
// @Description Query chaincode on peers
// @Param	body		body	chaincode.QueryArgs  true		"body for chaincode Description"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router /QueryCC [post]
func (u *ChaincodeController) QueryCC() {
	var req chaincode.QueryArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 301
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := chaincode.NewQueryAction(&req)
		if err != nil {
			fmt.Printf("Query Initialize error...")
			res["status"] = 307
			res["message"] = fmt.Sprintf("Query action error [%s]", err)
		} else {
			err := action.Query()
			if err != nil {
				res["status"] = 307
				res["message"] = fmt.Sprintf("Query execute error [%s]", err)
			} else {
				res["status"] = 200
				res["message"] = fmt.Sprintf("Query chaincode [%s] successfully", req.ChaincodeID)
			}
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
