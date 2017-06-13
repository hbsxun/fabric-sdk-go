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

// @Title AddAsset
// @Description create assets
// @Param	body		body 	models.Asset	true		"body for asset content"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router / [post]
func (u *AssetController) Post() {
	var asset models.Asset
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &asset)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	uid, err := models.AddAsset(asset)
	if err != nil {
		u.Data["json"] = err
	} else {
		u.Data["json"] = map[string]string{"uid": uid}
	}
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
