package cert

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/cert"
	"github.com/spf13/pflag"
)

// Operations about Enroll
type EnrollController struct {
	beego.Controller
}

// @Title Enroll
// @Description Get Key and Ecert
// @Param	body		body	cert.EnrollReq   true		"body for Ecert content"
// @Success 200 {string}
// @Failure 403 body is empty
// @router / [post]
func (u *EnrollController) Post() {
	var req cert.EnrollReq
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	flags := &pflag.FlagSet{}
	action, err := cert.NewEnrollAction(flags)
	if err != nil {
		fmt.Printf("Enroll Initialize error...")
	}
	resp, err := action.Enroll(req)
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = resp
	}
	fmt.Println(u.Data["json"])

	u.ServeJSON()
}
