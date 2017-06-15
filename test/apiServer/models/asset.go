package models

import (
	"encoding/json"

	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
)

type Asset struct {
	DocType string `json:"docType"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Desc    string `json:"desc"`
}

func AddAsset(a Asset) (string, error) {
	var req sdkIgn.Model
	//req.DocType = a.DocType
	req.Name = a.Name
	req.Owner = a.Owner
	req.Desc = a.Desc
	txId, err := setup.AddModel(&req)
	if err != nil {
		return "", err
	}
	return txId, nil
}

func GetAsset(assetId string) (a *Asset, err error) {
	assetInfo, err := setup.QueryModel(assetId)
	logger.Debug(assetInfo)
	if err != nil {
		return nil, err
	}
	var ass Asset
	err = json.Unmarshal([]byte(assetInfo), &ass)
	if err != nil {
		return nil, err
	}
	return &ass, nil
}

/*
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
*/
