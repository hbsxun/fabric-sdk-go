package handler

import (
	"encoding/json"

	"github.com/gislu/goSocket/server/utils"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
)

const (
	INIT       = "init"
	INVOKE     = "invoke"
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
	var req fabricCAClient.RegistrationRequest
	err = json.Unmarshal(contentJson, &req)
	if err != nil {
		utils.LogErr("msg.Content Unmarshal err %v", err)
		return nil
	}
	utils.Log("RegistrationRequest struct\n", req)

	admin := sdkIgn.NewMember()
	name, secret, err := admin.RegisterUser(&req)
	if err != nil {
		utils.LogErr("registerUser err ", err)
		return nil
	}

	utils.Logf("name [%s] secret [%s]", name, secret)
	var retMap = make(map[string]interface{})
	retMap["name"] = name
	retMap["secret"] = secret
	retJson, err := json.Marshal(retMap)
	if err != nil {
		utils.LogErr("retMap marshal to retJson err [%v]", err)
	}
	return retJson
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
}
