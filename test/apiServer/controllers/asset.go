package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/test/apiServer/models"
)

// Operations about Assets
type AssetController struct {
	beego.Controller
}

// @Title CreateAsset
// @Description create assets
// @Param	body		body 	models.Asset	true		"body for asset content"
// @Success 200 {int} models.Asset.Id
// @Failure 403 body is empty
// @router / [post]
func (u *AssetController) Post() {
	var asset models.Asset
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &asset)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	uid := models.AddAsset(asset)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Assets
// @Success 200 {object} models.Asset
// @router / [get]
func (u *AssetController) GetAll() {
	assets := models.GetAllAssets()
	fmt.Println("assets: ", assets)
	u.Data["json"] = assets
	u.ServeJSON()
}

// @Title Get
// @Description get asset by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Asset
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *AssetController) Get() {
	uid := u.GetString(":uid")
	fmt.Println("uid: ", uid)
	if uid != "" {
		asset, err := models.GetAsset(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = asset
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the asset
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Asset	true		"body for asset content"
// @Success 200 {object} models.Asset
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *AssetController) Put() {
	uid := u.GetString(":uid")
	if uid != "" {
		var asset models.Asset
		json.Unmarshal(u.Ctx.Input.RequestBody, &asset)
		uu, err := models.UpdateAsset(uid, &asset)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the asset
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *AssetController) Delete() {
	uid := u.GetString(":uid")
	models.DeleteAsset(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs asset into the system
// @Param	assetname		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 asset not exist
// @router /login [get]
func (u *AssetController) Login() {
	assetname := u.GetString("username")
	password := u.GetString("password")
	if models.Login(assetname, password) {
		u.Data["json"] = "login success"
	} else {
		u.Data["json"] = "asset not exist"
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in asset session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *AssetController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
