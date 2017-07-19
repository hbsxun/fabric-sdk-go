package query

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/query"
)

// Operations about QueryChannels
type QueryChannelsController struct {
	beego.Controller
}

// @Title QueryChannelsChaincodes
// @Description Query Chaincodes installed on the peerUrl
// @Param	peerUrl		path 	string		true	"The URL of peer to query"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /:peerUrl [get]
func (u *QueryChannelsController) Get() {
	peerUrl := u.GetString(":peerUrl")
	fmt.Println("peerUrl:", peerUrl)

	action, err := query.NewQueryChannelsAction(peerUrl)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		resp, err := action.Execute()
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = resp
		}
	}

	u.ServeJSON()
}
