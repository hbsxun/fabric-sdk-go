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
type info struct {
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
// @router /Adduser [post]
func (u *AssetController) Adduser() {
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

// @Title Add
// @Description Invoke chaincode on peers
// @Param	body		body	AssetApp.Args  true		"body for chaincode Description"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router /Add [post]
func (u *AssetController) Add() {
	var req AssetApp.Args
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := AssetApp.Add(&req)
	if err != nil {
		fmt.Printf("add  error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("add  successfully, txid = %s", resp)
	}

	u.ServeJSON()
}

// @Title Query
// @Description get  by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object}info
// @Failure 403 :id is empty
// @router /Query/:id [get]
func (u *AssetController) Queryasset() {
	id := u.GetString(":id")
	fmt.Println("id: ", id)
	if id != "" {
		resp, err := AssetApp.Query(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			var JSON info
			err = json.Unmarshal([]byte(resp), &JSON)
			u.Data["json"] = JSON
		}
	}
	u.ServeJSON()
}

// @Title Querybyowner
// @Description get  by owner
// @Param	owner		path 	string	true		"The key for staticblock"
// @Success 200 {string} list
// @Failure 403 :owner is empty
// @router /Querybyowner/:owner [get]
func (u *AssetController) Querybyowner() {
	owner := u.GetString(":owner")
	fmt.Println("owner: ", owner)
	if owner != "" {
		resp, err := AssetApp.QueryByOwner(owner)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = resp
		}
	}
	u.ServeJSON()
}

// @Title Gethistoryfor
// @Description get history for
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {string} history
// @Failure 403 :id is empty
// @router /gethistoryfor/:id [get]
func (u *AssetController) Gethistoryfor() {
	id := u.GetString(":id")
	fmt.Println("id: ", id)
	if id != "" {
		resp, err := AssetApp.GetHistoryFor(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = resp
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the , we can update "name", "owner", "price"
// @Param	body		body 	AssetApp.UpdateArgs	true		"body for update content"
// @Success 200 {string} txId
// @Failure 403 :id is empty
// @router /Update [put]
func (u *AssetController) Update() {
	var req AssetApp.UpdateArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := AssetApp.Update(&req)
	if err != nil {
		fmt.Printf("update  error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("update  successfully, txid = %s", resp)
	}

	u.ServeJSON()
}

// @Title Delete
// @Description delete the
// @Param	body		body 	AssetApp.DeleteArgs	true		"body for update content"
// @Success 200 {string} txId
// @Failure 403 :id is empty
// @router /Delete [put]
func (u *AssetController) Delete() {
	var req AssetApp.DeleteArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := AssetApp.Delete(&req)
	if err != nil {
		fmt.Printf("delete  error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("delete  successfully, txid = %s", resp)
	}

	u.ServeJSON()
}
