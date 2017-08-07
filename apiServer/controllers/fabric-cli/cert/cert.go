package cert

import (
	"encoding/json"
	"fmt"
	"strings"

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
func (u *CertController) Register() {
	var req cert.RegisterArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	res := make(map[string]interface{})
	if err != nil {
		res["status"] = 80400
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := cert.NewRegisterAction(&req)
		if err != nil {
			res["status"] = 80401
			res["message"] = fmt.Sprintf("NewRegisterAction failed [%s]", err.Error())
		}
		resp, err := action.Execute()
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("Register failed, err [%s]", err.Error())
		} else {
			res["status"] = 10200
			res["message"] = "Register Successfully, return an EnrollmentSecret"
			res["secret"] = resp
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}

// @Title Enroll
// @Description Get Key and Ecert
// @Param	body		body	cert.EnrollArgs   true		"body for Ecert content"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /Enroll [post]
func (u *CertController) Enroll() {
	var req cert.EnrollArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 80400
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := cert.NewEnrollAction(&req)
		if err != nil {
			res["status"] = 80401
			res["message"] = fmt.Sprintf("EnrollAction Initialize failed, err [%s]", err.Error())
		}
		resp, err := action.Execute()
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("Enroll failed, err [%s]", err.Error())
		} else {
			res["status"] = 80200
			res["message"] = "User Enroll Successfully, returns base64-encoded key and cert"
			keyCert := strings.Split(resp, ".")
			res["key"] = keyCert[0]
			res["cert"] = keyCert[1]
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
