package controllers

import (
	"encoding/json"
	"fmt"

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

// @Title GetAll
// @Description get all Registers
// @Success 200 {object} models.Register
// @router / [get]
func (u *RegisterController) GetAll() {
	u.ServeJSON()
}

// @Title Get
// @Description get identity by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Register
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *RegisterController) Get() {
	u.ServeJSON()
}

// @Title Update
// @Description update the identity
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Register	true		"body for identity content"
// @Success 200 {object} models.Register
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *RegisterController) Put() {
	u.ServeJSON()
}

// @Title Delete
// @Description delete the identity
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *RegisterController) Delete() {
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs identity into the system
// @Param	identityname		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 identity not exist
// @router /login [get]
func (u *RegisterController) Login() {
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in identity session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *RegisterController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
