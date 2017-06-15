package channel

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/channel"
)

// Operations about ChannelJoin
type ChannelJoinController struct {
	beego.Controller
}

// @Title ChannelJoin
// @Description ChannelJoin on peers
// @Param	body		body	channel.ChannelJoinArgs   true		"body for ChannelJoin Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router / [post]
func (u *ChannelJoinController) Post() {
	var req channel.ChannelJoinArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := channel.NewChannelJoinAction(&req)
	if err != nil {
		fmt.Printf("ChannelJoin Initialize error...")
	}
	err = action.Execute()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("Peer [%s] Joinchannel [%s] successful\n", req.PeerUrl, req.ChannelID)
	}

	u.ServeJSON()
}
