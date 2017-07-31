package query

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/query"
)

// @Title QueryInstalledChaincodes
// @Description Query Chaincodes installed on the peerUrl
// @Param	body		body 	query.QueryInstalledArgs		true	"body for Query Installed Chaincode Info"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /QueryInstalledChaincodes [post]
func (u *QueryController) QueryInstalledChaincodes() {
	var req query.QueryInstalledArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 301
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := query.NewqueryInstalledAction(&req)
		if err != nil {
			fmt.Printf("QueryInstalledChaincode Initialize error...")
			res["status"] = 403
			res["message"] = fmt.Sprintf("QueryInstalledChaincode action error [%s]", err)
		} else {
			err := action.Execute()
			if err != nil {
				res["status"] = 403
				res["message"] = fmt.Sprintf("QueryInstalledChaincode execute error [%s]", err)
			} else {
				res["status"] = 200
				res["message"] = fmt.Sprintf("query chaincodes installed on the [%s] successfully", req.PeerUrl)
			}
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
