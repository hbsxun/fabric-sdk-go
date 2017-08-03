package chaincode

import (
	"fmt"
	"testing"
)

func TestChaincode(t *testing.T) {
	peerUrl := "localhost:7051"
	channelID := "mychannel"
	chaincodeID := "example02_cc"
	chaincodePath := "github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/fixtures/src/github.com/example_cc"
	chaincodeVersion := "v0"

	fmt.Println("*****************************************************")
	installCC(&InstallCCArgs{
		//peerUrl,  //install on allPeers specified in config_test.yaml if not set peerUrl
		"",
		channelID,
		chaincodeID,
		chaincodePath,
		chaincodeVersion,
	}, t)

	fmt.Println("*****************************************************")
	instantiate(&InstantiateArgs{
		peerUrl,
		channelID,
		chaincodeID,
		chaincodePath,
		chaincodeVersion,
		[]string{"Init", "init", "a", "100", "b", "200"},
	}, t)

	fmt.Println("*****************************************************")
	query(&QueryArgs{
		peerUrl,
		channelID,
		chaincodeID,
		[]string{"invoke", "query", "a"},
	}, t)

	fmt.Println("*****************************************************")
	move(&InvokeArgs{
		peerUrl,
		channelID,
		chaincodeID,
		[]string{"invoke", "move", "a", "b", "200"},
	}, t)

	fmt.Println("*****************************************************")
	query(&QueryArgs{
		peerUrl,
		channelID,
		chaincodeID,
		[]string{"invoke", "query", "b"},
	}, t)

	fmt.Println("*****************************************************")
	chaincodeInfo(&ChaincodeInfoArgs{
		peerUrl,
		channelID,
		chaincodeID,
	}, t)
}

func installCC(args *InstallCCArgs, t *testing.T) {
	action, err := NewInstallAction(args)
	if err != nil {
		t.Fatal(err)
	}
	err = action.Execute()
	if err != nil {
		t.Fatal(err)
	}
}

func instantiate(args *InstantiateArgs, t *testing.T) {
	action, err := NewInstantiateAction(args)
	if err != nil {
		t.Fatal(err)
	}
	err = action.Execute()
	if err != nil {
		t.Fatal(err)
	}
}
func query(args *QueryArgs, t *testing.T) {
	action, err := NewQueryAction(args)
	if err != nil {
		t.Fatal(err)
	}
	_, err = action.Query()
	if err != nil {
		t.Fatal(err)
	}
}
func move(args *InvokeArgs, t *testing.T) {
	action, err := NewInvokeAction(args)
	if err != nil {
		t.Fatal(err)
	}
	err = action.Execute()
	if err != nil {
		t.Fatal(err)
	}
}
func chaincodeInfo(args *ChaincodeInfoArgs, t *testing.T) {
	action, err := NewChaincodeInfoAction(args)
	if err != nil {
		t.Fatal(err)
	}
	_, err = action.Execute()
	if err != nil {
		t.Fatal(err)
	}
}
