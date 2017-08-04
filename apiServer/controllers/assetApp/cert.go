package assetApp

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/assetApp/auth"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabricca"
)

// Operations about Register
type CertificateController struct {
	beego.Controller
}

// @Title Register
// @Description Get a OTP secret
// @Param	body		body	fabricca.RegisterArgs   true		"body for Secret content"
// @Success 200 {string} string
// @Failure 403 body is empty
// @router /Register [post]
func (u *CertificateController) Register() {
	var req fabricca.RegisterArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 307
		res["message"] = fmt.Sprintf("Unmarshal failed:%s", err.Error())
	} else {
		fmt.Println(req)
		action, err := fabricca.NewRegisterAction(&req)
		if err != nil {
			res["status"] = 322
			res["message"] = fmt.Sprintf("Register action failed:%s", err.Error())
		} else {
			resp, err := action.Execute()
			if err != nil {
				res["status"] = 323
				res["message"] = fmt.Sprintf("Register execute failed:%s", err.Error())
			} else {
				res["status"] = 200
				res["message"] = "Register in CA successfully"
				res["secret"] = resp
			}
		}
	}
	u.Data["json"] = res

	u.ServeJSON()
}

// @Title Enroll
// @Description Get Key and Ecert
// @Param	body		body	fabricca.EnrollArgs   true		"body for Ecert content"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /Enroll [post]
func (u *CertificateController) Enroll() {
	var req fabricca.EnrollArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		res["status"] = 307
		res["message"] = fmt.Sprintf("Unmarshal failed:%s", err.Error())
	} else {
		fmt.Println(req)
		action, err := fabricca.NewEnrollAction(&req)
		if err != nil {
			res["status"] = 320
			res["message"] = fmt.Sprintf("EnrollAction initialize failed, err [%s]", err.Error())
		} else {
			key, cert, err := action.Execute()
			if err != nil {
				res["status"] = 321
				res["message"] = "Failed to get Ecert and Key"
			} else {
				res["status"] = 200
				res["message"] = "Get Ecert and Key successfully"
				res["key"] = key
				res["cert"] = cert
			}
		}
	}
	//TODO the credentials of cert will be in secure device in the fulture
	_, userName := auth.GetIdAndName(u.Ctx.GetCookie("Bearer"))
	identity := auth.Identity{req.Name, req.Secret}
	//temporary use username to identify cert
	u.Ctx.SetCookie(userName, auth.Serialize(identity))

	u.Data["json"] = res
	u.ServeJSON()
}
