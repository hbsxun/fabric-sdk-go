package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/test/apiServer/models"
)

// Operations about Registers
type RegisterController struct {
	beego.Controller
}

// @Title Register
// @Description Get One-Time password for Ecert
// @Param	body		body 	models.RegisterRequest	true		"body for identity content"
// @Success 200 {object} models.RegisterResponse
// @Failure 403 body is empty
// @router / [post]
func (u *RegisterController) Post() {
	var registerReq models.RegisterRequest
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

// @Title Enroll
// @Description Get Key and Ecert
// @Param	body		body 	models.EnrollRequest	true		"body for Ecert content"
// @Success 200 {object} models.EnrollResponse
// @Failure 403 body is empty
// @router / [post]
func (u *EnrollController) Post() {
	var req models.EnrollRequest
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
