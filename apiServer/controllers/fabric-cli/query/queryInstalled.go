package query

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/query"
)

// Operations about QueryInstalled
type QueryInstalledController struct {
	beego.Controller
}

// @Title QueryInstalledChaincodes
// @Description Query Chaincodes installed on the peerUrl
// @Param	peerUrl		path 	string		true	"The URL of peer to query"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /:peerUrl [get]
func (u *QueryInstalledController) Get() {
	peerUrl := u.GetString(":peerUrl") //the peerUrl is empty, fixme, TODO
	fmt.Println("peerUrl:", peerUrl)
	if peerUrl == "" {
		peerUrl = "localhost:7051"
	}
	action, err := query.NewQueryInstalledAction(peerUrl)
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
