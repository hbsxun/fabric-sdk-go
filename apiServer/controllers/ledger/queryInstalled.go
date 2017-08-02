package ledger

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/query"
)

// @Title QueryInstalledChaincodes
// @Description Query chaincodes installed on the peerUrl
// @Param	peerUrl		path 	string		true	"The URL of peer to query"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /QueryInstalledChaincodes/:peerUrl [get]
func (u *LedgerController) QueryInstalledChaincodes() {
	peerUrl := u.GetString(":peerUrl") //the peerUrl is empty, fixme, TODO
	fmt.Println("peerUrl:", peerUrl)
	if peerUrl == "" {
		peerUrl = "localhost:7051"
	}
	res := make(map[string]interface{})
	action, err := query.NewQueryInstalledAction(peerUrl)
	if err != nil {
		res["status"] = 333
		res["message"] = fmt.Sprintf("NewQueryInstalledAction failed [%s]", err.Error())
	} else {
		resp, err := action.Execute()
		if err != nil {
			res["status"] = 333
			res["message"] = fmt.Sprintf("QueryInstalled execute failed [%s]", err.Error())
		} else {
			res["status"] = 200
			res["message"] = resp
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
