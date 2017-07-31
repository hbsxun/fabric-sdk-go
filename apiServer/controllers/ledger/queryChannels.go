package ledger

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/query"
)

// @Title QueryChannels
// @Description Query channels the peerUrl join
// @Param	peerUrl		path 	string		true	"The URL of peer to query"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /QueryChannels/:peerUrl [get]
func (u *LedgerController) QueryChannels() {
	peerUrl := u.GetString(":peerUrl")
	fmt.Println("peerUrl:", peerUrl)
	res := make(map[string]interface{})
	action, err := query.NewQueryChannelsAction(peerUrl)
	if err != nil {
		res["status"] = 332
		res["message"] = fmt.Sprintf("NewQueryChannelAction failed [%s]", err.Error())
	} else {
		resp, err := action.Execute()
		if err != nil {
			res["status"] = 332
			res["message"] = fmt.Sprintf("QueryChannel execute failed [%s]", err.Error())
		} else {
			res["status"] = 200
			res["message"] = resp
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
