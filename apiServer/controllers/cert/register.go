package cert

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabricca/cert"
)

// Operations about Register
type CertificateController struct {
	beego.Controller
}

// @Title Register
// @Description Get a OTP secret
// @Param	body		body	cert.RegisterArgs   true		"body for Secret content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /Register [post]
func (u *CertificateController) Register() {
	var req cert.RegisterArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 307
		res["secret"] = fmt.Sprintf("Unmarshal failed:%s", err.Error())
	} else {
		fmt.Println(req)
		action, err := cert.NewRegisterAction(&req)
		if err != nil {
			fmt.Printf("Register Initialize error...")
			res["status"] = 322
			res["secret"] = fmt.Sprintf("Register action failed:%s", err.Error())
		} else {
			resp, err := action.Execute()
			if err != nil {
				res["status"] = 323
				res["secret"] = fmt.Sprintf("Register execute failed:%s", err.Error())
			} else {
				res["status"] = 200
				res["secret"] = resp
			}
		}
	}
	u.Data["json"] = res
	fmt.Println(u.Data["json"])

	u.ServeJSON()
}
