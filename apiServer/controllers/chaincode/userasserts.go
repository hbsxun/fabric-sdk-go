package chaincode

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/chaincode"
)

// Operations about Invoke
type UserassertsController struct {
	beego.Controller
}
type Userinfo struct {
	ObjectType string `json:"doctype"`
	Id         string `json:"id"`
	Name       string `json:"name"`
}

//定义描述资产的数据结构
type Assertinfo struct {
	ObjectType string   `json:"doctype"`
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	User       Userinfo `json:"user"`
	Owner      string   `json:"owner"`
}

// @Title Adduser
// @Description Invoke chaincode on peers
// @Param	body		body	chaincode.UserArgs  true		"body for chaincode Description"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router /Adduser [post]
func (u *UserassertsController) Adduser() {
	var req chaincode.UserArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := chaincode.AddUser(&req)
	if err != nil {
		fmt.Printf("add user error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("add user successfully, txid = %s", resp)
	}

	u.ServeJSON()
}

// @Title Addassert
// @Description Invoke chaincode on peers
// @Param	body		body	chaincode.AssertArgs  true		"body for chaincode Description"
// @Success 200 {string} txId
// @Failure 403 body is empty
// @router /Addassert [post]
func (u *UserassertsController) Addassert() {
	var req chaincode.AssertArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := chaincode.AddAssert(&req)
	if err != nil {
		fmt.Printf("add assert error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("add assert successfully, txid = %s", resp)
	}

	u.ServeJSON()
}

// @Title Queryassert
// @Description get assert by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object}Assertinfo
// @Failure 403 :id is empty
// @router /Queryassert/:id [get]
func (u *UserassertsController) Queryassert() {
	assertid := u.GetString(":id")
	fmt.Println("assertid: ", assertid)
	if assertid != "" {
		resp, err := chaincode.QueryAssert(assertid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			var assertJSON Assertinfo
			err = json.Unmarshal([]byte(resp), &assertJSON)
			u.Data["json"] = assertJSON
		}
	}
	u.ServeJSON()
}

// @Title Queryassertbyowner
// @Description get assert by owner
// @Param	owner		path 	string	true		"The key for staticblock"
// @Success 200 {string} assertlist
// @Failure 403 :owner is empty
// @router /Querybyowner/:owner [get]
func (u *UserassertsController) Querybyowner() {
	owner := u.GetString(":owner")
	fmt.Println("owner: ", owner)
	if owner != "" {
		resp, err := chaincode.QueryAssertByOwner(owner)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = resp
		}
	}
	u.ServeJSON()
}

// @Title Gethistoryforassert
// @Description get history for assert
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {string} asserthistory
// @Failure 403 :id is empty
// @router /gethistoryforassert/:id [get]
func (u *UserassertsController) Gethistoryforassert() {
	id := u.GetString(":id")
	fmt.Println("assertid: ", id)
	if id != "" {
		resp, err := chaincode.GetHistoryForAssert(id)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = resp
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the assert, we can update "name", "owner", "price"
// @Param	body		body 	chaincode.UpdateArgs	true		"body for update content"
// @Success 200 {string} txId
// @Failure 403 :assertid is empty
// @router /Updateassert [put]
func (u *UserassertsController) Updateassert() {
	var req chaincode.UpdateArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := chaincode.UpdateAssert(&req)
	if err != nil {
		fmt.Printf("update assert error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("update assert successfully, txid = %s", resp)
	}

	u.ServeJSON()
}

// @Title Delete
// @Description delete the assert
// @Param	body		body 	chaincode.DeleteArgs	true		"body for update content"
// @Success 200 {string} txId
// @Failure 403 :assertid is empty
// @router /Deleteassert [put]
func (u *UserassertsController) Deleteassert() {
	var req chaincode.DeleteArgs
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &req)
	if err != nil {
		fmt.Printf("Unmarshal failed [%s]", err)
	}
	fmt.Println(req)
	resp, err := chaincode.DeleteAssert(&req)
	if err != nil {
		fmt.Printf("delete assert error...")
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("delete assert successfully, txid = %s", resp)
	}

	u.ServeJSON()
}
