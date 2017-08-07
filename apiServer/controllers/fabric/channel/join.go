package channel

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/channel"
)

// @Title ChannelJoin
// @Description ChannelJoin on peers
// @Param	body		body	channel.ChannelJoinArgs   true		"body for ChannelJoin Description"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /JoinChannel [post]
func (u *ChannelController) JoinChannel() {
	var req channel.ChannelJoinArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 80401
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := channel.NewChannelJoinAction(&req)
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("NewChannelJoinAction failed [%s]", err)
		} else {
			err = action.Execute()
			if err != nil {
				res["status"] = 80403
				res["message"] = fmt.Sprintf("ChannelJoin Execute failed [%s]", err)
			} else {
				res["status"] = 80200
				if req.ChannelID == "" && req.PeerUrl == "" {
					res["message"] = fmt.Sprintf("All peers Joinchannel [mychannel] successfully")
				} else if req.ChannelID == "" && req.PeerUrl != "" {
					res["message"] = fmt.Sprintf("Peer [%s] Joinchannel [mychannel] successfully", req.PeerUrl)
				} else if req.ChannelID != "" && req.PeerUrl == "" {
					res["message"] = fmt.Sprintf("All peers Joinchannel [%s] successfully", req.ChannelID)
				} else {
					res["message"] = fmt.Sprintf("Peer [%s] Joinchannel [%s] successfully", req.PeerUrl, req.ChannelID)
				}
			}
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}
