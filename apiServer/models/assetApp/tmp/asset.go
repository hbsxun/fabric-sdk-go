package AssetApp

import (
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/chaincode"
)

type AssetArgs struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Owner    string `json:"owner"`
}
type UpdateArgs struct {
	Id            string `json:"id"`
	Updatetag     string `json:"updatetag"`
	Updatecontent string `json:"updatecontent"`
	Userid        string `json:"ueserid"`
	Username      string `json:"username"`
}

type DeleteArgs struct {
	Id       string `json:"id"`
	Userid   string `json:"ueserid"`
	Username string `json:"username"`
}

func AddAsset(asset *AssetArgs) (string, error) {
	var args []string
	args = append(args, "AddAsset")
	//args = append(args, model.DocType)
	args = append(args, asset.Id)
	args = append(args, asset.Name)
	args = append(args, asset.Price)
	args = append(args, asset.Userid)
	args = append(args, asset.Username)
	args = append(args, asset.Owner)
	action, err := chaincode.NewInvokeAction(&chaincode.InvokeArgs{"testchannel", "asset_cc", args})
	if err != nil {
		return "", err
	}
	txid, err := action.Execute()
	if err != nil {
		return "", err
	}
	return txid, nil
}

func QueryAsset(id string) (string, error) {
	var args []string
	args = append(args, "QueryAsset")
	args = append(args, id)
	action, err := chaincode.NewQueryAction(&chaincode.QueryArgs{"testchannel", "asset_cc", args})
	if err != nil {
		return "", err
	}
	result, err := action.Execute()
	if err != nil {
		return "", err
	}
	return result, nil
}

func QueryAssetByOwner(owner string) (string, error) {
	var args []string
	args = append(args, "queryAssetByOwner")
	args = append(args, owner)
	action, err := chaincode.NewQueryAction(&chaincode.QueryArgs{"testchannel", "asset_cc", args})
	if err != nil {
		return "", err
	}
	result, err := action.Execute()
	if err != nil {
		return "", err
	}
	return result, nil
}

func GetHistoryForAsset(id string) (string, error) {
	var args []string
	args = append(args, "getHistoryForAsset")
	args = append(args, id)
	action, err := chaincode.NewQueryAction(&chaincode.QueryArgs{"testchannel", "asset_cc", args})
	if err != nil {
		return "", err
	}
	result, err := action.Execute()
	if err != nil {
		return "", err
	}
	return result, nil
}

func UpdateAsset(updatearg *UpdateArgs) (string, error) {
	var args []string
	args = append(args, "UpdateAsset")
	args = append(args, updatearg.Id)
	args = append(args, updatearg.Updatetag)
	args = append(args, updatearg.Updatecontent)
	args = append(args, updatearg.Userid)
	args = append(args, updatearg.Username)
	action, err := chaincode.NewInvokeAction(&chaincode.InvokeArgs{"testchannel", "asset_cc", args})
	if err != nil {
		return "", err
	}
	txid, err := action.Execute()
	if err != nil {
		return "", err
	}
	return txid, nil
}

func DeleteAsset(deletearg *DeleteArgs) (string, error) {
	var args []string
	args = append(args, "DeleteAsset")
	args = append(args, deletearg.Id)
	args = append(args, deletearg.Userid)
	args = append(args, deletearg.Username)
	action, err := chaincode.NewInvokeAction(&chaincode.InvokeArgs{"testchannel", "asset_cc", args})
	if err != nil {
		return "", err
	}
	txid, err := action.Execute()
	if err != nil {
		return "", err
	}
	return txid, nil
}
