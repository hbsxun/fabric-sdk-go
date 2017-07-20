package user

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	user.User	true		"body for user content"
// @Success 200 {int} user.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var ur user.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &ur)
	uid, err := user.AddUser(&ur)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = uid
	}
	u.ServeJSON()
}

// @Title Get
// @Description get user by name
// @Param	name		path 	string	true		"The key for staticblock"
// @Success 200 {object} user.User
// @Failure 403 :name is empty
// @router /:name [get]
func (u *UserController) Get() {
	uid := u.GetString(":name")
	fmt.Println("name: ", uid)
	if uid != "" {
		user, err := user.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	user.User	true		"body for user content"
// @Success 200 {object} user.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	var ur user.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &ur)
	err := user.UpdateUser(&ur)
	if err != nil {
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = "update success"
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	body		body 	user.Secret	true		"The username for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var ss user.Secret
	json.Unmarshal(u.Ctx.Input.RequestBody, &ss)

	fmt.Println("Secret:", ss)

	u, err := user.Login(&ss)
	if err != nil {
		u.Data["json"] = fmt.Errorf("login failed, err [%s]", err.Error())
	} else {
		u.Data["json"] = u
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
