package AssetApp

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/AssetApp"
)

// Operations about Invoke
type AssetController struct {
	beego.Controller
}
type Userinfo struct {
	ObjectType string `json:"doctype"`
	Id         string `json:"id"`
	Name       string `json:"name"`
}

//定义描述资产的数据结构
type Assetinfo struct {
	ObjectType string   `json:"doctype"`
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	User       Userinfo `json:"user"`
	Owner      string   `json:"owner"`
}

// @Title Adduser
// @Description Invoke chaincode on peers
// @Param	body		body	AssetApp.UserArgs  true		"body for chaincode Description"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router /AddUser [post]
func (u *AssetController) AddUser() {
	var req AssetApp.UserArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := AssetApp.AddUser(&req)
	if err != nil {
		fmt.Printf("add user error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("add user successfully, txid = %s", resp)
	}

	u.ServeJSON()
}

// @Title AddAsset
// @Description Invoke chaincode on peers
// @Param	body		body	AssetApp.AssetArgs  true		"body for chaincode Description"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router /AddAsset [post]
func (u *AssetController) AddAsset() {
	var req AssetApp.AssetArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := AssetApp.AddAsset(&req)
	if err != nil {
		fmt.Printf("add asset error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("add asset successfully, txid = %s", resp)
	}

	u.ServeJSON()
}

// @Title QueryAsset
// @Description get asset by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object}Assetinfo
// @Failure 403 :id is empty
// @router /QueryAsset/:id [get]
func (u *AssetController) QueryAsset() {
	assetid := u.GetString(":id")
	fmt.Println("assetid: ", assetid)
	if assetid != "" {
		resp, err := AssetApp.QueryAsset(assetid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			var assetJSON Assetinfo
			err = json.Unmarshal([]byte(resp), &assetJSON)
			u.Data["json"] = assetJSON
		}
	}
	u.ServeJSON()
}

// @Title Queryassetbyowner
// @Description get asset by owner
// @Param	owner		path 	string	true		"The key for staticblock"
// @Success 200 {string} assetlist
// @Failure 403 :owner is empty
// @router /QueryAssetByOwner/:owner [get]
func (u *AssetController) QueryAssetByOwner() {
	owner := u.GetString(":owner")
	fmt.Println("owner: ", owner)
	if owner != "" {
		resp, err := AssetApp.QueryAssetByOwner(owner)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = resp
		}
	}
	u.ServeJSON()
}

// @Title Gethistoryforasset
// @Description get history for asset
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {string} assethistory
// @Failure 403 :id is empty
// @router /GetAssetTradeHistory/:id [get]
func (u *AssetController) GetAssetTradeHistory() {
	id := u.GetString(":id")
	fmt.Println("assetid: ", id)
	if id != "" {
		resp, err := AssetApp.GetHistoryForAsset(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = resp
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the asset, we can update "name", "owner", "price"
// @Param	body		body 	AssetApp.UpdateArgs	true		"body for update content"
// @Success 200 {string} txId
// @Failure 403 :assetid is empty
// @router /UpdateAsset [put]
func (u *AssetController) UpdateAsset() {
	var req AssetApp.UpdateArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := AssetApp.UpdateAsset(&req)
	if err != nil {
		fmt.Printf("update asset error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("update asset successfully, txid = %s", resp)
	}

	u.ServeJSON()
}

// @Title Delete
// @Description delete the asset
// @Param	body		body 	AssetApp.DeleteArgs	true		"body for update content"
// @Success 200 {string} txId
// @Failure 403 :assetid is empty
// @router /DeleteAsset [put]
func (u *AssetController) DeleteAsset() {
	var req AssetApp.DeleteArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := AssetApp.DeleteAsset(&req)
	if err != nil {
		fmt.Printf("delete asset error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("delete asset successfully, txid = %s", resp)
	}

	u.ServeJSON()
}
