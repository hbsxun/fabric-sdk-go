package assetApp

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/assetApp"
)

// Operations about Initialize
type InitializeController struct {
	beego.Controller
}

// @Title Initialize
// @Description create channel, join channel, and instantiate chaincode
// @Success 200 {string} install and instantiate chaincode successfully
// @Failure 403 body is empty
// @router / [post]
func (u *InitializeController) Post() {
	err := assetApp.Initialize()
	if err != nil {
		fmt.Printf("Initialize error:%s", err.Error())
		u.Data["json"] = err.Error()
	} else {
		u.Data["json"] = fmt.Sprintf("Initialize successfully")
	}
	u.ServeJSON()
}
