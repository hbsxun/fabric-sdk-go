package handler

import (
	"encoding/json"
	"os"

	"github.com/gislu/goSocket/server/utils"
	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
)

var setup *sdkIgn.BaseSetupImpl
var prefix = os.Getenv("GOPATH") + "/src/github.com/hyperledger/fabric-sdk-go/test"

const (
	ADDASSET   = "addAsset"
	QUERYASSET = "queryAsset"
)

type AddAssetController struct {
}

func (this *AddAssetController) Excute(msg utils.Msg) []byte {
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

//QueryAssetController
type QueryAssetController struct {
}

func (this *QueryAssetController) Excute(msg utils.Msg) []byte {
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

	modelInfo, err := setup.QueryModel(req.Name)
	if err != nil {
		utils.LogErr("QueryModel failed ", err)
	}
	utils.Logf("modelInfo [%s]\n", modelInfo)

	return []byte(modelInfo)
}
func init() {
	//add AddAssetController
	var addAsset AddAssetController
	utils.Route(func(entry utils.Msg) bool {
		if entry.Meta["meta"] == ADDASSET {
			return true
		}
		return false
	}, &addAsset)
	//add QueryAssetController
	var queryAsset QueryAssetController
	utils.Route(func(entry utils.Msg) bool {
		if entry.Meta["meta"] == QUERYASSET {
			return true
		}
		return false
	}, &queryAsset)

	//get BaseSetupImpl instance and initialize asset/model chaincode
	setup = sdkIgn.NewBaseSetupImpl(prefix)
	err := setup.InstallAndInstantiateModelCC()
	if err != nil {
		utils.LogErr("InstallAndInstantiateModelCC failed ", err)
		os.Exit(-1)
	}
}
