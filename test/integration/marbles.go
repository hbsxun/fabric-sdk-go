package integration

// InstallAndInstantiateExampleCC ..
func (setup *BaseSetupImpl) InstallAndInstantiateMarblesCC() error {

	chainCodePath := "github.com/marbles_cc"
	chainCodeVersion := "v0"

	if setup.ChainCodeID == "" {
		//setup.ChainCodeID = fcUtil.GenerateRandomID()
		setup.ChainCodeID = "marbles_cc"
	}

	if err := setup.InstallCC(setup.ChainCodeID, chainCodePath, chainCodeVersion, nil); err != nil {
		return err
	}

	var args []string
	args = append(args, "init")
	args = append(args, "520")

	return setup.InstantiateCC(setup.ChainCodeID, setup.ChainID, chainCodePath, chainCodeVersion, args)
}

func (setup *BaseSetupImpl) read() (string, error) {

	var args []string
	args = append(args, "read")
	args = append(args, "selftest")
	return setup.Query(setup.ChainID, setup.ChainCodeID, args)
}

/*
// Querymarbles ...
// Addmarbles ...
func (setup *BaseSetupImpl) Addmarbles(model *Model) (string, error) {

	var args []string
	args = append(args, "initmarbles")
	//args = append(args, marbles.DocType)
	args = append(args, marbles.Owner)
	args = append(args, marbles.Name)
	args = append(args, marbles.Desc)

	transientDataMap := make(map[string][]byte)
	transientDataMap["result"] = []byte("Transient data in add marbles...")

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
*/
