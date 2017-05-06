package integration

import (
	"fmt"
	"time"

	fcUtil "github.com/hyperledger/fabric-sdk-go/fabric-client/helpers"

	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
)

// InstallAndInstantiateExampleCC ..
func (setup *BaseSetupImpl) InstallAndInstantiateExampleCC() error {

	chainCodePath := "github.com/example_cc"
	chainCodeVersion := "v0"

	if setup.ChainCodeID == "" {
		setup.ChainCodeID = fcUtil.GenerateRandomID()
	}

	if err := setup.InstallCC(setup.ChainCodeID, chainCodePath, chainCodeVersion, nil); err != nil {
		return err
	}

	var args []string
	args = append(args, "init")
	args = append(args, "a")
	args = append(args, "100")
	args = append(args, "b")
	args = append(args, "200")

	return setup.InstantiateCC(setup.ChainCodeID, setup.ChainID, chainCodePath, chainCodeVersion, args)
}

// QueryAsset ...
func (setup *BaseSetupImpl) QueryAsset() (string, error) {

	var args []string
	args = append(args, "invoke")
	args = append(args, "query")
	args = append(args, "b")
	return setup.Query(setup.ChainID, setup.ChainCodeID, args)
}

// MoveFunds ...
func (setup *BaseSetupImpl) MoveFunds() (string, error) {

	var args []string
	args = append(args, "invoke")
	args = append(args, "move")
	args = append(args, "a")
	args = append(args, "b")
	args = append(args, "1")

	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in move funds...")

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
