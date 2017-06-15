package cert

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/cert"
)

// Operations about Register
type RegisterController struct {
	beego.Controller
}

// @Title Register
// @Description Get a OTP secret
// @Param	body		body	cert.RegisterArgs   true		"body for Secret content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router / [post]
func (u *RegisterController) Post() {
	var req cert.RegisterArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := cert.NewRegisterAction(&req)
	if err != nil {
		fmt.Printf("Register Initialize error...")
	}
	resp, err := action.Execute()
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = resp
	}
	fmt.Println(u.Data["json"])

	u.ServeJSON()
}
