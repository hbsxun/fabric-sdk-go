package assetApp

import (
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/chaincode"
)

const (
	defaultChannelID   = "mychannel"
	defaultChaincodeID = "model_cc"
	defaultPeerUrl     = "localhost:7051"
)

//Asset = Model in model_cc chaincode
type AddModelArgs struct {
	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Price string `json:"price"`
	Owner string `json:"owner"`
}
type TransferModelArgs struct {
	Name     string `json:"name"`
	NewOwner string `json:"newOwner"`
}

//AddModel returns a transaction id
func AddModel(model *AddModelArgs) error {
	args := []string{model.Name, model.Desc, model.Price, model.Owner}
	invokeArgs := &chaincode.InvokeArgs{
		PeerUrl:     defaultPeerUrl,
		ChannelID:   defaultChannelID,
		ChaincodeID: defaultChaincodeID,
		Args:        append([]string{"addModel"}, args...),
	}
	invokeAction, err := chaincode.NewInvokeAction(invokeArgs)
	if err != nil {
		appLogger.Debugf("NewInvokeAction err [%s]\n", err.Error())
		return err
	}
	err = invokeAction.Execute()
	if err != nil {
		appLogger.Debugf("invokeAction err [%s]\n", err.Error())
		return err
	}
	return nil
}

func QueryModel(modelName string) (string, error) {
	queryArgs := &chaincode.QueryArgs{
		PeerUrl:     defaultPeerUrl,
		ChannelID:   defaultChannelID,
		ChaincodeID: defaultChaincodeID,
		Args:        append([]string{"queryModel"}, modelName),
	}
	queryAction, err := chaincode.NewQueryAction(queryArgs)
	if err != nil {
		appLogger.Debugf("NewQueryAction err [%s]\n", err.Error())
		return "", err
	}
	res, err := queryAction.Query()
	if err != nil {
		appLogger.Debugf("queryAction err [%s]\n", err.Error())
		return "", err
	}
	return res, err
}

func TransferModel(model *TransferModelArgs) error {
	args := []string{model.Name, model.NewOwner}
	invokeArgs := &chaincode.InvokeArgs{
		PeerUrl:     defaultPeerUrl,
		ChannelID:   defaultChannelID,
		ChaincodeID: defaultChaincodeID,
		Args:        append([]string{"transferModel"}, args...),
	}
	invokeAction, err := chaincode.NewInvokeAction(invokeArgs)
	if err != nil {
		appLogger.Debugf("NewInvokeAction err [%s]\n", err.Error())
		return err
	}
	err = invokeAction.Execute()
	if err != nil {
		appLogger.Debugf("invokeAction err [%s]\n", err.Error())
		return err
	}
	return nil

}

func GetHistoryForModel(modelName string) (string, error) {
	queryArgs := &chaincode.QueryArgs{
		PeerUrl:     defaultPeerUrl,
		ChannelID:   defaultChannelID,
		ChaincodeID: defaultChaincodeID,
		Args:        append([]string{"getHistoryForModel"}, modelName),
	}
	queryAction, err := chaincode.NewQueryAction(queryArgs)
	if err != nil {
		appLogger.Debugf("NewQueryAction err [%s]\n", err.Error())
		return "", err
	}
	res, err := queryAction.Query()
	if err != nil {
		appLogger.Debugf("queryAction err [%s]\n", err.Error())
		return "", err
	}
	return res, err

}

func QueryModelsByOwner(owner string) (string, error) {
	queryArgs := &chaincode.QueryArgs{
		PeerUrl:     defaultPeerUrl,
		ChannelID:   defaultChannelID,
		ChaincodeID: defaultChaincodeID,
		Args:        append([]string{"queryModelsByOwner"}, owner),
	}
	queryAction, err := chaincode.NewQueryAction(queryArgs)
	if err != nil {
		appLogger.Debugf("NewQueryAction err [%s]\n", err.Error())
		return "", err
	}
	res, err := queryAction.Query()
	if err != nil {
		appLogger.Debugf("queryAction err [%s]\n", err.Error())
		return "", err
	}
	return res, err
}

func DelModel(modelName string) error {
	args := []string{modelName}
	invokeArgs := &chaincode.InvokeArgs{
		ChannelID:   defaultChannelID,
		ChaincodeID: defaultChaincodeID,
		Args:        append([]string{"delModel"}, args...),
	}
	invokeAction, err := chaincode.NewInvokeAction(invokeArgs)
	if err != nil {
		appLogger.Debugf("NewInvokeAction err [%s]\n", err.Error())
		return err
	}
	err = invokeAction.Execute()
	if err != nil {
		appLogger.Debugf("invokeAction err [%s]\n", err.Error())
		return err
	}
	return nil

}

/*
//rich query
func QueryModels() { }
*/
