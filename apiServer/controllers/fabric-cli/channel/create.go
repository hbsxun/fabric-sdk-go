package channel

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/channel"
)

// Operations about ChannelCreate
type ChannelController struct {
	beego.Controller
}

// @Title ChannelCreate
// @Description ChannelCreate on peers
// @Param	body		body	channel.ChannelCreateArgs   true		"body for ChannelCreate Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /CreateChannel [post]
func (u *ChannelController) Post() {
	var req channel.ChannelCreateArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 80401
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		fmt.Println(len(req.ChannelID), len(req.OrdererUrl), len(req.TxFile))
		action, err := channel.NewChannelCreateAction(&req)
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("NewCreateChannelAction failed [%s]", err)
		} else {
			err = action.Execute()
			if err != nil {
				res["status"] = 80403
				res["message"] = fmt.Sprintf("CreateChannel Execute error [%s]", err)
			} else {
				res["status"] = 80200
				if req.ChannelID == "" {
					res["message"] = fmt.Sprintf("Channel create [mychannel] successfully")
				} else {
					res["message"] = fmt.Sprintf("Channel create [%s] successfully", req.ChannelID)
				}
			}
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}
