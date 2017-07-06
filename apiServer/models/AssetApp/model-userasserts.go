package AssetApp

import (
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/chaincode"
)

type UserArgs struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type AssertArgs struct {
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

func AddUser(user *UserArgs) (string, error) {
	var args []string
	args = append(args, "AddUser")
	//args = append(args, model.DocType)
	args = append(args, user.Id)
	args = append(args, user.Name)
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

func AddAssert(assert *AssertArgs) (string, error) {
	var args []string
	args = append(args, "AddAssert")
	//args = append(args, model.DocType)
	args = append(args, assert.Id)
	args = append(args, assert.Name)
	args = append(args, assert.Price)
	args = append(args, assert.Userid)
	args = append(args, assert.Username)
	args = append(args, assert.Owner)
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

func QueryAssert(id string) (string, error) {
	var args []string
	args = append(args, "QueryAssert")
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

func QueryAssertByOwner(owner string) (string, error) {
	var args []string
	args = append(args, "queryAssertByOwner")
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

func GetHistoryForAssert(id string) (string, error) {
	var args []string
	args = append(args, "getHistoryForAssert")
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

func UpdateAssert(updatearg *UpdateArgs) (string, error) {
	var args []string
	args = append(args, "UpdateAssert")
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

func DeleteAssert(deletearg *DeleteArgs) (string, error) {
	var args []string
	args = append(args, "DeleteAssert")
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
