package channel

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/channel"
)

// Operations about ChannelCreate
type ChannelCreateController struct {
	beego.Controller
}

// @Title ChannelCreate
// @Description ChannelCreate on peers
// @Param	body		body	channel.ChannelCreateArgs   true		"body for ChannelCreate Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router / [post]
func (u *ChannelCreateController) Post() {
	var req channel.ChannelCreateArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := channel.NewChannelCreateAction(&req)
	if err != nil {
		fmt.Printf("ChannelCreate Initialize error...")
	}
	err = action.Execute()
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("Channel create  [%s] successful\n", req.ChannelID)
	}

	u.ServeJSON()
}
