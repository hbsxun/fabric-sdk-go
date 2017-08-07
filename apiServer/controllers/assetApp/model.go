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
		res["status"] = 80401
		res["message"] = fmt.Sprintf("AddModelArgs Unmarshal failed [%s]", err)
	} else {
		fmt.Println(req)
		err = assetApp.AddModel(&req)
		if err != nil {
			res["status"] = 80403
			res["message"] = fmt.Sprintf("AddModel failed[%s]", err.Error())
		} else {
			res["status"] = 80200
			res["message"] = fmt.Sprintf("AddModel successfully")
		}
	}
	fmt.Println(res)

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
			res["message"] = fmt.Sprintf("QueryModel failed[%s]", err.Error())
			res["status"] = 80403
		} else {
			var modelJSON assetApp.AddModelArgs
			err = json.Unmarshal([]byte(resp), &modelJSON)
			if err != nil {
				res["message"] = fmt.Sprintf("QueryModelRes Unmarshal failed[%s]", err.Error())
				res["status"] = 80401
			} else {
				res["status"] = 80200
				res["message"] = modelJSON
			}
		}
	}
	fmt.Println(res)

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
		res["status"] = 80401
		res["message"] = fmt.Sprintf("TransferModelArgs Unmarshal failed[%s]", err.Error())
	}
	fmt.Println(req)
	err = assetApp.TransferModel(&req)
	if err != nil {
		res["status"] = 8043
		res["message"] = fmt.Sprintf("TransferModel failed[%s]", err.Error())
	} else {
		res["status"] = 80200
		res["message"] = fmt.Sprintf("TransferModel successfully")
	}
	fmt.Println(res)

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
			res["status"] = 80403
			res["message"] = fmt.Sprintf("QueryModelsByOwner failed[%s]", err.Error())
		} else {
			res["status"] = 80200
			res["message"] = resp
		}
	}
	fmt.Println(res)

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
			res["status"] = 80403
			res["message"] = fmt.Sprintf("GetHistoryForModel failed[%s]", err.Error())
		} else {
			res["status"] = 80200
			res["message"] = resp
		}
	}
	fmt.Println(res)

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
		err := assetApp.DelModel(name)
		if err != nil {
			res["status"] = 80403
			res["message"] = fmt.Sprintf("DeleteModel failed[%s]", err.Error())
		} else {
			res["status"] = 80200
			res["message"] = fmt.Sprintf("DeleteModel successfully")
		}
	}
	fmt.Println(res)

	u.Data["json"] = res
	u.ServeJSON()
}
