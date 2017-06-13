package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/test/apiServer/models"
)

// Operations about Ledgers
type LedgerController struct {
	beego.Controller
}

// @Title GetTransaction
// @Description Get transaction from ledger
// @Param	uid		path 	string	true		"The key for transaction"
// @Success 200 {object} models.Transaction
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *LedgerController) Get() {
	uid := u.GetString(":uid")
	fmt.Println("uid: ", uid)
	if uid != "" {
		var txInfo *models.Transaction
		txInfo, err := models.GetTx(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = txInfo
		}
	} else {
		u.Data["json"] = "No txId specified"
	}
	u.ServeJSON()
}
