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
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/spf13/pflag"
)

type InstallCCArgs struct {
	ChaincodeName string `json:"chaincodeName"`
	//ChaincodeID      string `json:"chaincodeId"`
	ChaincodeVersion string `json:"chaincodeVersion"`
}

type installAction struct {
	common.ActionImpl
}

func NewInstallAction(flags *pflag.FlagSet) (*installAction, error) {
	action := &installAction{}
	err := action.Initialize(flags)
	return action, err
}

func (action *installAction) Invoke() error {
	fmt.Printf("Installing chaincode %s on peers:\n", common.ChaincodeID)
	for _, peer := range action.Peers() {
		fmt.Printf("-- %s\n", peer.GetURL())
	}

	err := installChaincode(action.Client(), action.Peers(), common.ChaincodeID, common.ChaincodePath, common.ChaincodeVersion)
	if err != nil {
		return err
	}

	fmt.Printf("...successfuly installed chaincode %s on peers:\n", common.ChaincodeID)
	for _, peer := range action.Peers() {
		fmt.Printf("-- %s\n", peer.GetURL())
	}

	return nil
}

func installChaincode(client fabricClient.Client, targets []fabricClient.Peer, chaincodeID string, chaincodePath string, chaincodeVersion string) error {
	var errors []error

	transactionProposalResponse, _, err := client.InstallChaincode(chaincodeID, chaincodePath, chaincodeVersion, nil, targets)
	if err != nil {
		return fmt.Errorf("InstallChaincode returned error: %v", err)
	}

	ccIDVersion := chaincodeID + "." + chaincodeVersion

	for _, v := range transactionProposalResponse {
		if v.Err != nil {
			if strings.Contains(v.Err.Error(), ccIDVersion+" exists") {
				// Ignore
				common.Logger.Infof("Chaincode %s already installed on peer: %s.\n", ccIDVersion, v.Endorser)
			} else {
				errors = append(errors, fmt.Errorf("installCC returned error from peer %s: %v", v.Endorser, v.Err))
			}
		} else {
			common.Logger.Infof("...successfuly installed chaincode %s on peer %s.\n", ccIDVersion, v.Endorser)
		}
	}

	if len(errors) > 0 {
		common.Logger.Warningf("Errors returned from InstallCC: %v\n", errors)
		return errors[0]
	}

	return nil
}
