package cert

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/cert"
)

// @Title Enroll
// @Description Get Key and Ecert
// @Param	body		body	cert.EnrollArgs   true		"body for Ecert content"
// @Success 200 {string}
// @Failure 403 body is empty
// @router /Enroll [post]
func (u *CertificateController) Enroll() {
	var req cert.EnrollArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 307
		res["key"] = fmt.Sprintf("Unmarshal failed:%s", err.Error())
		res["cert"] = fmt.Sprintf("Unmarshal failed:%s", err.Error())
	} else {
		fmt.Println(req)
		action, err := cert.NewEnrollAction(&req)
		if err != nil {
			res["status"] = 320
			res["key"] = fmt.Sprintf("Enroll action failed:%s", err.Error())
			res["cert"] = fmt.Sprintf("Enroll action failed:%s", err.Error())
		} else {
			key, cert, err := action.Execute()
			if err != nil {
				res["status"] = 321
				res["key"] = fmt.Sprintf("Enroll execute failed:%s", err.Error())
				res["cert"] = fmt.Sprintf("Enroll execute failed:%s", err.Error())
			} else {
				res["status"] = 200
				res["key"] = key
				res["cert"] = cert
			}
		}
	}
	u.Data["json"] = res
	fmt.Println(u.Data["json"])

	u.ServeJSON()
}
