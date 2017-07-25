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
	id, ok := impl.Register(&ur)
	res := make(map[string]interface{})
	if !ok {
		res["status"] = 301
		res["message"] = "register failed"
	} else {
		res["status"] = 200
		res["message"] = "register successfully"
		res["userId"] = id
	}
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
	json.Unmarshal(u.Ctx.Input.RequestBody, &ss)
	impl := assetApp.UserManagerImpl{}
	signedToekn, err := impl.Login(&ss)
	fmt.Println("User Secret:", ss)
	fmt.Println("Cookie Token:", signedToekn)
	res := make(map[string]interface{})
	if err != nil {
		res["status"] = 302
		res["message"] = "Login failed"
	} else {
		//for authorization
		u.Ctx.SetCookie("Bearer", signedToekn)
		res["status"] = 200
		res["message"] = "Login successfully"
	}
	u.Data["json"] = res
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
	res := make(map[string]interface{})
	if err != nil {
		res["status"] = 303
		res["message"] = fmt.Sprintf("update failed:%s", err.Error())
	} else {
		res["status"] = 200
		res["message"] = "User update successfully"
	}
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
	if name != "" {
		impl := assetApp.UserManagerImpl{}
		userInfo, err := impl.GetUserInfoByName(name)
		fmt.Println("userInfo: ", userInfo)
		res := make(map[string]interface{})
		if err != nil {
			res["status"] = 304
			res["message"] = fmt.Sprintf("get user by name failed:%s", err.Error())
		} else {
			res["status"] = 200
			res["message"] = userInfo
		}
		u.Data["json"] = res
	}

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
	if id != "" {
		res := make(map[string]interface{})
		userId, err := strconv.Atoi(id)
		if err != nil {
			res["status"] = 304
			res["message"] = fmt.Sprintf("atoi failed:%s", err.Error())
		} else {
			impl := assetApp.UserManagerImpl{}
			userInfo, err := impl.GetUserInfoById(userId)
			fmt.Println("userInfo: ", userInfo)
			if err != nil {
				res["status"] = 304
				res["message"] = fmt.Sprintf("get user by id failed:%s", err.Error())
			} else {
				res["status"] = 200
				res["message"] = userInfo
			}
		}
		u.Data["json"] = res
	}

	u.ServeJSON()
}

// @Title Logout
// @Description logout
// @Success 200 {string}logout successfully
// @Failure 403 :logout failed
// @router /logout [get]
func (u *UserManageController) Logout() {
	u.Ctx.SetCookie("Bearer", "xxxx")
	res := make(map[string]interface{})
	res["status"] = 200
	res["message"] = "logout successfully"

	u.Data["json"] = res

	u.ServeJSON()
}
