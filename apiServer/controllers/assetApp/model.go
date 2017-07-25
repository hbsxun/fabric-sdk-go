package assetApp

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	assetApp "github.com/hyperledger/fabric-sdk-go/apiServer/models/assetApp"
)

// Operations about Invoke
type AssetController struct {
	beego.Controller
}

// @Title AddModel
// @Description Invoke chaincode on peers
// @Param	body		body	assetApp.AddModelArgs  true		"body for chaincode Description"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router /AddModel [post]
func (u *AssetController) AddModel() {
	var req assetApp.AddModelArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 307
		res["message"] = "Unmarshal failed"
	} else {
		fmt.Println(req)
		resp, err := assetApp.AddModel(&req)
		if err != nil {
			res["message"] = fmt.Sprintf("Add model error:%s", err.Error())
			res["status"] = 308
		} else {
			res["status"] = 200
			res["message"] = fmt.Sprintf("Add model successfully")
			res["txid"] = resp
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}

// @Title QueryModel
// @Description get model by name
// @Param	ModelName		path 	string	true		"The key for staticblock"
// @Success 200 {object}assetApp.AddModelArgs
// @Failure 403 :ModelName is empty
// @router /QueryModel/:ModelName [get]
func (u *AssetController) QueryModel() {
	name := u.GetString(":ModelName")
	fmt.Println("name: ", name)
	res := make(map[string]interface{})
	if name != "" {
		resp, err := assetApp.QueryModel(name)
		if err != nil {
			res["message"] = fmt.Sprintf("query model error:%s", err.Error())
			res["status"] = 309
		} else {
			var modelJSON assetApp.AddModelArgs
			err = json.Unmarshal([]byte(resp), &modelJSON)
			res["status"] = 200
			res["message"] = modelJSON
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}

// @Title TransferModel
// @Description Invoke chaincode on peers
// @Param	body		body 	assetApp.TransferModelArgs	true		"body for chaincode content"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router /TransferModel [put]
func (u *AssetController) TransferModel() {
	var req assetApp.TransferModelArgs
	res := make(map[string]interface{})
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
		res["status"] = 307
		res["message"] = "Unmarshal failed"
	}
	fmt.Println(req)
	resp, err := assetApp.TransferModel(&req)
	if err != nil {
		fmt.Printf("Transfer model error:%s", err.Error())
		res["status"] = 308
		res["message"] = fmt.Sprintf("transfer model error:%s", err.Error())
	} else {
		res["status"] = 200
		res["message"] = fmt.Sprintf("Transfer model successfully")
		res["txid"] = resp
	}
	u.Data["json"] = res
	u.ServeJSON()
}

// @Title QueryModelsByOwner
// @Description query models by owner
// @Param	owner		path 	string	true		"The key for staticblock"
// @Success 200 {string} ModelList
// @Failure 403 :owner is empty
// @router /QueryModelsByOwner/:owner [get]
func (u *AssetController) QueryModelsByOwner() {
	owner := u.GetString(":owner")
	res := make(map[string]interface{})
	fmt.Println("owner:", owner)
	if owner != "" {
		resp, err := assetApp.QueryModelsByOwner(owner)
		if err != nil {
			res["status"] = 309
			res["message"] = fmt.Sprintf("query model by owner error:%s", err.Error())
		} else {
			res["status"] = 200
			res["message"] = resp
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}

// @Title GetHistoryForModel
// @Description get history for model
// @Param	ModelName		path 	string	true		"The key for staticblock"
// @Success 200 {string} ModelHistory
// @Failure 403 :ModelName is empty
// @router /GetHistoryForModel/:ModelName [get]
func (u *AssetController) GetHistoryForModel() {
	name := u.GetString(":ModelName")
	fmt.Println("name: ", name)
	res := make(map[string]interface{})
	if name != "" {
		resp, err := assetApp.GetHistoryForModel(name)
		if err != nil {
			res["status"] = 309
			res["message"] = fmt.Sprintf("get history for model error:%s", err.Error())
		} else {
			res["status"] = 200
			res["message"] = resp
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}

// @Title DeleteModel
// @Description delete model
// @Param	ModelName		path 	string	true		"The key for staticblock"
// @Success 200 {string} txId
// @Failure 403 :ModelName is empty
// @router /DeleteModel/:ModelName [put]
func (u *AssetController) DeleteModel() {
	name := u.GetString(":ModelName")
	fmt.Println("name: ", name)
	res := make(map[string]interface{})
	if name != "" {
		resp, err := assetApp.DelModel(name)
		if err != nil {
			res["status"] = 308
			res["message"] = fmt.Sprintf("Delete model error:%s", err.Error())
		} else {
			res["status"] = 200
			res["message"] = fmt.Sprintf("Delete model successfully")
			res["txid"] = resp
		}
	}
	u.Data["json"] = res
	u.ServeJSON()
}
