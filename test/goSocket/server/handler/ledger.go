package handler

import (
	"encoding/json"

	"github.com/gislu/goSocket/server/utils"
	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
)

var ledger *sdkIgn.Ledger

const (
	QUERYBLOCK       = "queryBlock"       //on ledger
	QUERYTRANSACTION = "queryTransaction" //on ledger
)

type QueryBlockController struct {
}

func (this *QueryBlockController) Excute(msg utils.Msg) []byte {
	utils.Log("*********************************************")
	utils.Log(msg.Content)

	contentJson, err := json.Marshal(msg.Content)
	if err != nil {
		utils.LogErr("msg.Content marshal err %v", err)
		return nil
	}

	var req sdkIgn.Model
	err = json.Unmarshal(contentJson, &req)
	if err != nil {
		utils.LogErr("msg.Content Unmarshal err %v", err)
		return nil
	}
	utils.Log("Request Model struct\n", req)

	txId, err := setup.AddModel(&req)
	if err != nil {
		utils.LogErr("AddModel failed ", err)
	}

	utils.Logf("txId [%s]\n", txId)
	var retMap = make(map[string]interface{})
	retMap["txId"] = txId
	retJson, err := json.Marshal(retMap)
	if err != nil {
		utils.LogErr("retMap marshal to retJson err [%v]", err)
	}
	return retJson
}

func init() {
	//add AddAssetController
	var queryBlock QueryBlockController
	utils.Route(func(entry utils.Msg) bool {
		if entry.Meta["meta"] == QUERYBLOCK {
			return true
		}
		return false
	}, &queryBlock)

	//get BaseSetupImpl instance and initialize asset/model chaincode
	ledger = sdkIgn.NewLedger(setup.Chain)
}
