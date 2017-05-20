package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("Models  Asset")

var setup *sdkIgn.BaseSetupImpl
var prefix = os.Getenv("GOPATH") + "/src/github.com/hyperledger/fabric-sdk-go/test"

type Asset struct {
	DocType string `json:"docType"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Desc    string `json:"desc"`
}

func AddAsset(a Asset) string {
	var req sdkIgn.Model
	//req.DocType = a.DocType
	req.Name = a.Name
	req.Owner = a.Owner
	req.Desc = a.Desc
	txId, err := setup.AddModel(&req)
	if err != nil {
		log.Errorf("AddAsset failed [%s]", err)
	}
	return txId
}

func GetAsset(assetId string) (a *Asset, err error) {
	assetInfo, err := setup.QueryModel(assetId)
	fmt.Println(assetInfo)
	if err != nil {
		log.Errorf("QueryAsset failed [%s]", err)
		return nil, err
	}
	var ass Asset
	err = json.Unmarshal([]byte(assetInfo), &ass)
	if err != nil {
		log.Errorf("Marshal Asset failed [%s]", err)
		return nil, err
	}
	return &ass, nil
}

func GetAllAssets() map[string]*Asset {
	return nil
}

func UpdateAsset(uid string, uu *Asset) (a *Asset, err error) {
	return nil, errors.New("Asset Not Exist")
}

func Login(name, password string) bool {
	return true
}

func DeleteAsset(uid string) {
}

func init() {
	setup = sdkIgn.NewBaseSetupImpl(prefix)
	err := setup.InstallAndInstantiateModelCC()
	if err != nil {
		log.Errorf("InstallAndInstantiateModelCC failed ", err)
		os.Exit(-1)
	}
}
