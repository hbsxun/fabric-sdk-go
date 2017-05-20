package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-ca/api"

	"github.com/astaxie/beego"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
	"github.com/hyperledger/fabric-sdk-go/test/apiServer/models"
)

// Operations about Registers
type RegisterController struct {
	beego.Controller
}

// @Title CreateRegister
// @Description create identitys
// @Param	body		body 	models.Register	true		"body for identity content"
// @Success 200 {int} models.Register.Id
// @Failure 403 body is empty
// @router / [post]
func (u *RegisterController) Post() {
	var registerReq fabricCAClient.RegistrationRequest
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &registerReq)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	resp, err := models.Register(&registerReq)
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = resp
	}
	fmt.Println(u.Data["json"])

	u.ServeJSON()
}

// Operations about Registers
type EnrollController struct {
	beego.Controller
}

// @Title CreateRegister
// @Description create identitys
// @Param	body		body 	models.Register	true		"body for identity content"
// @Success 200 {int} models.Register.Id
// @Failure 403 body is empty
// @router / [post]
func (u *EnrollController) Post() {
	var req api.EnrollmentRequest
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	resp, err := models.Enroll(&req)
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = resp
	}
	fmt.Println(u.Data["json"])

	u.ServeJSON()
}
