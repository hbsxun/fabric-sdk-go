package integration

import (
	"fmt"
	"time"

	fcUtil "github.com/hyperledger/fabric-sdk-go/fabric-client/helpers"

	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
)

type Model struct {
	//DocType string `json:"docType"`
	Owner string `json:"owner"`
	Name  string `json:"name"`
	Desc  string `json:"desc"`
}

// InstallAndInstantiateExampleCC ..
func (setup *BaseSetupImpl) InstallAndInstantiateModelCC() error {

	chainCodePath := "github.com/model_cc"
	chainCodeVersion := "v0"

	if setup.ChainCodeID == "" {
		setup.ChainCodeID = fcUtil.GenerateRandomID()
	}

	if err := setup.InstallCC(setup.ChainCodeID, chainCodePath, chainCodeVersion, nil); err != nil {
		return err
	}

	var args []string
	args = append(args, "init")

	return setup.InstantiateCC(setup.ChainCodeID, setup.ChainID, chainCodePath, chainCodeVersion, args)
}

// QueryModel ...
func (setup *BaseSetupImpl) QueryModel(modelId string) (string, error) {

	var args []string
	args = append(args, "readModel")
	args = append(args, modelId)
	return setup.Query(setup.ChainID, setup.ChainCodeID, args)
}

// TransferModel ...
func (setup *BaseSetupImpl) TransferModel() (string, error) {

	var args []string
	args = append(args, "transferModel")
	args = append(args, "M1")
	args = append(args, "Bob")

	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in transfer model...")

	transactionProposalResponse, txID, err := fcUtil.CreateAndSendTransactionProposal(setup.Chain, setup.ChainCodeID, setup.ChainID, args, []fabricClient.Peer{setup.Chain.GetPrimaryPeer()}, transientDataMap)
	if err != nil {
		return "", fmt.Errorf("CreateAndSendTransactionProposal return error: %v", err)
	}
	// Register for commit event
	done, fail := fcUtil.RegisterTxEvent(txID, setup.EventHub)

	txResponse, err := fcUtil.CreateAndSendTransaction(setup.Chain, transactionProposalResponse)
	if err != nil {
		return "", fmt.Errorf("CreateAndSendTransaction return error: %v", err)
	}
	fmt.Println(txResponse)
	select {
	case <-done:
	case <-fail:
		return "", fmt.Errorf("invoke Error received from eventhub for txid(%s) error(%v)", txID, fail)
	case <-time.After(time.Second * 30):
		return "", fmt.Errorf("invoke Didn't receive block event for txid(%s)", txID)
	}
	return txID, nil

}

// AddModel ...
func (setup *BaseSetupImpl) AddModel(model *Model) (string, error) {

	var args []string
	args = append(args, "initModel")
	//args = append(args, model.DocType)
	args = append(args, model.Owner)
	args = append(args, model.Name)
	args = append(args, model.Desc)

	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in add model...")

	transactionProposalResponse, txID, err := fcUtil.CreateAndSendTransactionProposal(setup.Chain, setup.ChainCodeID, setup.ChainID, args, []fabricClient.Peer{setup.Chain.GetPrimaryPeer()}, transientDataMap)
	if err != nil {
		return "", fmt.Errorf("CreateAndSendTransactionProposal return error: %v", err)
	}
	// Register for commit event
	done, fail := fcUtil.RegisterTxEvent(txID, setup.EventHub)

	txResponse, err := fcUtil.CreateAndSendTransaction(setup.Chain, transactionProposalResponse)
	if err != nil {
		return "", fmt.Errorf("CreateAndSendTransaction return error: %v", err)
	}
	fmt.Println(txResponse)
	select {
	case <-done:
	case <-fail:
		return "", fmt.Errorf("invoke Error received from eventhub for txid(%s) error(%v)", txID, fail)
	case <-time.After(time.Second * 30):
		return "", fmt.Errorf("invoke Didn't receive block event for txid(%s)", txID)
	}
	return txID, nil

}

// QueryModelByOwner by Owner ...
func (setup *BaseSetupImpl) QueryModelByOwner(owner string) (string, error) {

	var args []string
	args = append(args, "queryModelsByOwner")
	//args = append(args, "Alice")
	args = append(args, owner)
	return setup.Query(setup.ChainID, setup.ChainCodeID, args)
}

//GetHistoryForModel ...
func (setup *BaseSetupImpl) GetHistoryForModel() (string, error) {

	var args []string
	args = append(args, "getHistoryForModel")
	args = append(args, "M1")
	return setup.Query(setup.ChainID, setup.ChainCodeID, args)
}
