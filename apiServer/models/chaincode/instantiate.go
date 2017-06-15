/*
Copyright SecureKey Technologies Inc. All Rights Reserved.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at


      http://www.apache.org/licenses/LICENSE-2.0


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package chaincode

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/hyperledger/fabric-sdk-go/fabric-client/events"
	"github.com/spf13/pflag"
)

/*
flags.String(common.PeerFlag, "", "The URL of the peer on which to instantiate the chaincode, e.g. localhost:7051")
	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, common.ChannelID, "The channel ID")
	flags.StringVar(&common.ChaincodeID, common.ChaincodeIDFlag, "", "The chaincode ID")
	flags.StringVar(&common.ChaincodePath, common.ChaincodePathFlag, "", "The chaincode path")
	flags.StringVar(&common.ChaincodeVersion, common.ChaincodeVersionFlag, common.ChaincodeVersion, "The chaincode version")
	flags.StringVar(&common.Args, common.ArgsFlag, common.Args, "The args in JSON format. Example: {\"Args\":\"invoke\",\"arg1\",\"arg2\"}")
*/

type InstantiateArgs struct {
	//OrdererUrl       string           `json:"orderUrl"`
	//PeerUrl          string           `json:"peerUrl"`
	//ChannelID        string           `json:"channelID"`
	//ChaincodePath    string           `json:"chaincodePath"`
	//ChaincodeVersion string           `json:"chaincodeVersion"`
	ChaincodeID string   `json:"chaincodeId"`
	Args        []string `json:"args"`
}

type instantiateAction struct {
	common.ActionImpl
}

func NewInstantiateAction(args *InstantiateArgs) (*instantiateAction, error) {
	action, flags := &instantiateAction{}, &pflag.FlagSet{}

	flags.StringVar(&common.OrdererURL, common.OrdererFlag, "localhost:7050", "")
	flags.StringVar(&common.PeerURL, common.PeerFlag, "localhost:7051", "")
	flags.StringVar(&common.ChaincodePath, common.ChaincodePathFlag, common.CCPathPrefix+args.ChaincodeID, "")
	flags.StringVar(&common.ChaincodeID, common.ChaincodeIDFlag, args.ChaincodeID, "")
	flags.StringVar(&common.Args, common.ArgsFlag, common.GetMarshalArgs(common.ArgStruct{args.Args}), "")

	err := action.Initialize(flags)
	if len(action.Peers()) == 0 {
		return nil, fmt.Errorf("a peer must be specified")
	}

	fmt.Printf("%s\n%s\n%s\n%s\n%s\n", common.OrdererURL, common.PeerURL, common.ChaincodePath, common.ChannelID, common.Args)

	return action, err
}

func (action *instantiateAction) Execute() (string, error) {
	chain, err := action.NewChain() //just search testchannel, don't create
	//chain, err := fcUtil.GetChain(action.Client(), common.ChannelID) //create channel if not found
	//join the channel first, note!!
	if err != nil {
		return "", fmt.Errorf("Error initializing chain: %v", err)
	}

	argBytes := []byte(common.Args)
	args := &common.ArgStruct{}
	err = json.Unmarshal(argBytes, args)
	if err != nil {
		return "", fmt.Errorf("Error unmarshaling JSON arg string: %v", err)
	}

	fmt.Printf("instantiating chaincode on peer: %v\n", action.Peers()[0].GetURL())

	txId, err := instantiateChaincode(chain, action.EventHub(), common.ChannelID, common.ChaincodeID, common.ChaincodeVersion, common.ChaincodePath, args.Args)
	if err != nil {
		return "", fmt.Errorf("Error instantiating chaincode: %v", err)
	}

	fmt.Printf("...successfuly instantiated chaincode %s on channel %s.\n", common.ChaincodeID, common.ChannelID)

	return txId, nil
}

func instantiateChaincode(chain fabricClient.Chain, eventHub events.EventHub, chainID string, chaincodeID string, chainCodeVersion string, chainCodePath string, args []string) (string, error) {
	txId, err := instantiateCC(chain, eventHub, chainID, chaincodeID, chainCodeVersion, chainCodePath, args)
	if err != nil {
		if strings.Contains(err.Error(), "Chaincode exists "+chaincodeID) {
			// Ignore
			common.Logger.Infof("Chaincode %s already instantiated.", chaincodeID)
			return "", err
		}
		return "", fmt.Errorf("instantiateCC returned error: %v", err)
	}

	return txId, nil
}

func instantiateCC(chain fabricClient.Chain, eventHub events.EventHub, chainID string, chainCodeID string, chainCodeVersion string, chainCodePath string, args []string) (string, error) {
	transactionProposalResponse, txID, err := chain.SendInstantiateProposal(chainCodeID, chainID, args, chainCodePath, chainCodeVersion, []fabricClient.Peer{chain.GetPrimaryPeer()})
	if err != nil {
		return "", fmt.Errorf("SendInstantiateProposal return error: %v", err)
	}

	for _, v := range transactionProposalResponse {
		if v.Err != nil {
			return "", fmt.Errorf("SendInstantiateProposal Endorser %s return error: %v", v.Endorser, v.Err)
		}
		fmt.Printf("SendInstantiateProposal Endorser '%s' return ProposalResponse status:%v\n", v.Endorser, v.Status)
	}

	tx, err := chain.CreateTransaction(transactionProposalResponse)
	if err != nil {
		return "", fmt.Errorf("CreateTransaction return error: %v", err)

	}
	transactionResponse, err := chain.SendTransaction(tx)
	if err != nil {
		return "", fmt.Errorf("SendTransaction return error: %v", err)

	}
	for _, v := range transactionResponse {
		if v.Err != nil {
			return "", fmt.Errorf("Orderer %s return error: %v", v.Orderer, v.Err)
		}
	}
	done := make(chan bool)
	fail := make(chan error)

	eventHub.RegisterTxEvent(txID, func(txID string, err error) {
		if err != nil {
			fail <- err
		} else {
			fmt.Printf("instantiateCC receive success event for txid(%s)\n", txID)
			done <- true
		}

	})

	select {
	case <-done:
	case <-fail:
		return "", fmt.Errorf("instantiateCC Error received from eventhub for txid(%s) error(%v)", txID, fail)
	case <-time.After(time.Second * 60):
		return "", fmt.Errorf("timed out waiting to receive block event for txid(%s)", txID)
	}
	return txID, nil

}
