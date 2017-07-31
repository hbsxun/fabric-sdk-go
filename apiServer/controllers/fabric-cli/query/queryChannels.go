package query

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/query"
)

// @Title QueryChannels
// @Description Query Channels
// @Param	body		body 	query.QueryChannelsArgs		true	"body for Query Channel"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /QueryChannels [post]
func (u *QueryController) QueryChannels() {
	var req query.QueryChannelsArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 301
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := query.NewQueryChannelsAction(&req)
		if err != nil {
			fmt.Printf("QueryChannel Initialize error...")
			res["status"] = 401
			res["message"] = fmt.Sprintf("QueryChannel action error [%s]", err)
		} else {
			err := action.Execute()
			if err != nil {
				res["status"] = 401
				res["message"] = fmt.Sprintf("QueryChannel execute error [%s]", err)
			} else {
				res["status"] = 200
				res["message"] = "query channel successfully"
			}
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
