package fabricca

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabricca"
)

// Operations about Register
type CertController struct {
	beego.Controller
}

// @Title Register
// @Description Get a OTP secret
// @Param	body		body	fabricca.RegisterArgs   true		"body for Secret content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /Register [post]
func (u *CertController) Register() {
	var req fabricca.RegisterArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	res := make(map[string]interface{})
	if err != nil {
		res["status"] = 80400
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := fabricca.NewRegisterAction(&req)
		if err != nil {
			res["status"] = 80401
			res["message"] = fmt.Sprintf("NewRegisterAction failed [%s]", err.Error())
		}
		resp, err := action.Execute()
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("Register Execute failed [%s]", err.Error())
		} else {
			res["status"] = 80200
			res["message"] = "Register Successfully, return an EnrollmentSecret [secret]"
			res["secret"] = resp
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}

// @Title Enroll
// @Description Get Key and Ecert
// @Param	body		body	fabricca.EnrollArgs   true		"body for Ecert content"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /Enroll [post]
func (u *CertController) Enroll() {
	var req fabricca.EnrollArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 80400
		res["message"] = fmt.Sprintf("Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		action, err := fabricca.NewEnrollAction(&req)
		if err != nil {
			res["status"] = 80401
			res["message"] = fmt.Sprintf("NewEnrollAction failed, err [%s]", err.Error())
		}
		resp, err := action.Execute()
		if err != nil {
			res["status"] = 80402
			res["message"] = fmt.Sprintf("Enroll Execute failed, err [%s]", err.Error())
		} else {
			res["status"] = 80200
			res["message"] = "User Enroll Successfully, returns base64(key) and base64(cert) contacted by '.'"
			res["key.cert"] = resp
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}
