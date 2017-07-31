package cert

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/cert"
)

// Operations about Register
type CertController struct {
	beego.Controller
}

// @Title Register
// @Description Get a OTP secret
// @Param	body		body	cert.RegisterArgs   true		"body for Secret content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /Register [post]
func (u *CertController) Post() {
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
	res := make(map[string]interface{})
	if err != nil {
		res["status"] = 80401
		res["message"] = err.Error()
	} else {
		res["status"] = 80200
		res["message"] = "[MSP] Reigster user Successfully"
		res["secret"] = resp
	}
	u.Data["json"] = res
	u.ServeJSON()
}

// @Title Enroll
// @Description Get Key and Ecert
// @Param	body		body	cert.EnrollArgs   true		"body for Ecert content"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /Enroll [post]
func (u *CertController) Post() {
	var req cert.EnrollArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	action, err := cert.NewEnrollAction(&req)
	if err != nil {
		fmt.Printf("Enroll Initialize error...")
	}
	resp, err := action.Execute()
	res := make(map[string]interface{})
	if err != nil {
		res["status"] = 80401
		res["message"] = err.Error()
	} else {
		res["status"] = 80200
		res["message"] = "[MSP] Reigster user Successfully"
		res["secret"] = resp
	}
	u.Data["json"] = res

	u.ServeJSON()
}
