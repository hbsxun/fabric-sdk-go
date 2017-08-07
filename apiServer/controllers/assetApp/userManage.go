package assetApp

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	id, err := impl.Register(&ur)
	res := make(map[string]interface{})
	if err != nil {
		res["status"] = 80403
		res["message"] = fmt.Sprintf("UserManagerImpl Register failed [%s]", err.Error())
	} else {
		res["status"] = 80200
		res["message"] = "Register successfully, return a userId"
		res["userId"] = id
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	body		body 	user.Secret	true		"body for user login"
// @Success 200 {string} login successfully
// @Failure 403 user not exist or login failed
// @router /userLogin [post]
func (u *UserManageController) Login() {
	var ss user.Secret
	res := make(map[string]interface{})

	json.Unmarshal(u.Ctx.Input.RequestBody, &ss)
	fmt.Println("User Secret:", ss)

	impl := assetApp.UserManagerImpl{}
	signedToekn, err := impl.Login(&ss)
	if err != nil {
		res["status"] = 80403
		res["message"] = fmt.Sprintf("UserManagerImpl Login failed[%s]", err.Error())
	} else {
		fmt.Println("Cookie Token:", signedToekn)
		//for authorization
		u.Ctx.SetCookie("Bearer", signedToekn)
		res["status"] = 80200
		res["message"] = "Login successfully"
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}

// @Title UpdateInfo
// @Description update the user
// @Param	body		body 	user.UpdateUserArgs	true		"body for user content"
// @Success 200 {string} update successfully
// @Failure 403 :uid is not int
// @router /updateUser [put]
func (u *UserManageController) UpdateInfo() {
	var ur user.UpdateUserArgs
	res := make(map[string]interface{})

	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ur)
	if err != nil {
		res["status"] = 80401
		res["message"] = err.Error()
	}
	fmt.Println("updateUser:", ur)

	impl := assetApp.UserManagerImpl{}
	err = impl.UpdateInfo(&ur)
	if err != nil {
		res["status"] = 80403
		res["message"] = fmt.Sprintf("UpdateUser failed[%s]", err.Error())
	} else {
		res["status"] = 80200
		res["message"] = "updateUser successfully"
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}

// @Title UpdatePasswd
// @Description update password
// @Param	name		path 	string	true		"The name of user"
// @Param	oldPassword		path 	string	true		"The old password of user"
// @Param	newPassword		path 	string	true		"The new password of user"
// @Success 200 {string}
// @Failure 403 :name is empty
// @router /UpdatePasswd/:name/:oldPassword/:newPassword [put]
func (u *UserManageController) UpdatePasswd() {
	name, oldPwd, newPwd := u.GetString(":name"), u.GetString(":oldPassword"), u.GetString(":newPassword")
	fmt.Println("userName: ", name, "oldPwd: ", oldPwd, "newPwd: ", newPwd)

	res := make(map[string]interface{})
	if name != "" && oldPwd != "" && newPwd != "" {
		impl := assetApp.UserManagerImpl{}
		err := impl.UpdatePwd(name, oldPwd, newPwd)
		if err != nil {
			res["status"] = 80403
			res["message"] = fmt.Sprintf("UpdatePwd [%s]", err.Error())
		} else {
			res["status"] = 80200
			res["message"] = "UpdatePwd successfully"
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}

// @Title GetUserByName
// @Description get user by username
// @Param	userName		path 	string	true		"The key for staticblock"
// @Success 200 {object}user.User
// @Failure 403 :userName is empty
// @router /getUserByName/:userName [get]
func (u *UserManageController) GetUserByName() {
	name := u.GetString(":userName")
	fmt.Println("userName: ", name)

	res := make(map[string]interface{})
	if name != "" {
		impl := assetApp.UserManagerImpl{}
		userInfo, err := impl.GetUserInfoByName(name)
		fmt.Println("userInfo: ", userInfo)
		if err != nil {
			res["status"] = 80403
			res["message"] = fmt.Sprintf("GetUserByName failed[%s]", err.Error())
		} else {
			res["status"] = 80200
			res["message"] = userInfo
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}

// @Title GetUserById
// @Description get user by userid
// @Param	userId		path 	string	true		"The key for staticblock"
// @Success 200 {object}user.User
// @Failure 403 :userId is empty
// @router /getUserById/:userId [get]
func (u *UserManageController) GetUserById() {
	id := u.GetString(":userId")
	fmt.Println("userId: ", id)

	res := make(map[string]interface{})
	if id != "" {
		userId, err := strconv.Atoi(id)
		if err != nil {
			res["status"] = 80401
			res["message"] = fmt.Sprintf("Format strconv string->int failed [%s]", err.Error())
		} else {
			impl := assetApp.UserManagerImpl{}
			userInfo, err := impl.GetUserInfoById(userId)
			fmt.Println("userInfo: ", userInfo)
			if err != nil {
				res["status"] = 80403
				res["message"] = fmt.Sprintf("GetUserById failed [%s]", err.Error())
			} else {
				res["status"] = 80200
				res["message"] = userInfo
			}
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}

// @Title Logout
// @Description logout
// @Success 200 {string}logout successfully
// @Failure 403 :logout failed
// @router /logout [get]
func (u *UserManageController) Logout() {
	u.Ctx.SetCookie("Bearer", "The field has been reset.")

	res := make(map[string]interface{})
	res["status"] = 80200
	res["message"] = "Logout successfully"
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}

// @Title VerifyUser
// @Description verify user
// @Param	body		body 	user.Secret	true		"body for user login"
// @Success 200 {string} verify successfully
// @Failure 403 user not exist or login failed
// @router /VerifyUser [post]
func (u *UserManageController) VerifyUser() {
	var ss user.Secret
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &ss)
	if err != nil {
		res["status"] = 80401
		res["message"] = fmt.Sprintf("user.Secret [%s]", err.Error())
	}
	fmt.Println("User Secret:", ss)

	impl := assetApp.UserManagerImpl{}
	err = impl.VerifyUser(&ss)
	if err != nil {
		res["status"] = 80403
		res["message"] = fmt.Sprintf("VerifyUser failed [%s]", err.Error())
	} else {
		res["status"] = 80200
		res["message"] = "Verify successfully"
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}
