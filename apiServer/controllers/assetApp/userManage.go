package assetApp

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/assetApp"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
)

// Operations about Users
type UserManageController struct {
	beego.Controller
}

// @Title Register
// @Description register user
// @Param	body		body 	user.User	true		"body for user content"
// @Success 200 {int} user.User.Id
// @Failure 403 body is empty
// @router /addUser [post]
func (u *UserManageController) Register() {
	var ur user.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &ur)
	impl := assetApp.UserManagerImpl{}
	uid, ok := impl.Register(&ur)
	if !ok {
		u.Data["json"] = "register failed"
	} else {
		u.Data["json"] = fmt.Sprintf("Register user successfully, uid:%d", uid)
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login successfully
// @Failure 403 user not exist
// @router /userLogin [get]
func (u *UserManageController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	fmt.Println("name:", username, "passwd:", password)
	impl := assetApp.UserManagerImpl{}
	signedToekn, err := impl.Login(username, password)
	if err != nil {
		u.Data["json"] = "login failed!"
	} else {
		//for authorization
		u.Ctx.SetCookie("Bearer", signedToekn)
		tokenMap := make(map[string]interface{})
		tokenMap["status"] = 200
		tokenMap["message"] = "Login successfully"
		u.Data["json"] = tokenMap
	}
	u.ServeJSON()
}

// @Title UpdateInfo
// @Description update the user
// @Param	body		body 	user.User	true		"body for user content"
// @Success 200 {string} update successfully
// @Failure 403 :uid is not int
// @router /updateUser [put]
func (u *UserManageController) UpdateInfo() {
	var ur user.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &ur)
	fmt.Println("updateUser:", ur)
	impl := assetApp.UserManagerImpl{}
	err := impl.UpdateInfo(&ur)
	if err != nil {
		u.Data["json"] = fmt.Sprintf("update failed:%s", err.Error())
	} else {
		u.Data["json"] = "update successfully"
	}
	u.ServeJSON()
}
