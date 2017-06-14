package query

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/query"
	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	"github.com/spf13/pflag"
)

// Operations about QueryInstalled
type QueryInstalledController struct {
	beego.Controller
}

// @Title QueryInstalledChaincodes
// @Description Query Chaincodes installed on the peerURL
// @Param	peerURL		path 	string		true	"The URL of peer to query"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /:peerURL [get]
func (u *QueryInstalledController) Get() {
	peerURL := u.GetString("peerURL") //the peerURL is empty, fixme, TODO
	fmt.Println("peerURL:", peerURL)
	if peerURL == "" {
		peerURL = "localhost:7051"
	}
	flags := &pflag.FlagSet{}
	flags.StringVar(&query.PeerURL, common.PeerFlag, peerURL, "")
	action, err := query.NewQueryInstalledAction(flags)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		ccs, err := action.Query()
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = ccs
		}
	}

	u.ServeJSON()
}
