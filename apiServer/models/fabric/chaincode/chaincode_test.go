package chaincode

import (
	"testing"
	"time"
)

func TestChaincode(t *testing.T) {
	install(t)
	time.Sleep(time.Second * 5)
	initialize(t)
}

func install(t *testing.T) {
	//install chaincode on peers
	installAction, err := NewInstallAction(&InstallCCArgs{
		ChaincodeName:    "model_cc",
		ChaincodeVersion: "v0",
	})
	if err != nil {
		t.Fatal(err)
	}
	if err = installAction.Execute(); err != nil {
		t.Fatal(err)
	}

}

func initialize(t *testing.T) {
	//initialize chaincode on primary peer
	initAction, err := NewInstantiateAction(&InstantiateArgs{
		ChaincodeID: "model_cc",
		Args:        nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, err = initAction.Execute(); err != nil {
		t.Fatal(err)
	}
}
