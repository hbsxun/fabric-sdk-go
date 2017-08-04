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

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/spf13/pflag"
)

/*
	flags.String(common.PeerFlag, "", "The URL of the peer to connect to, e.g. localhost:7051")
	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, common.ChannelID, "The channel ID")
	flags.StringVar(&common.ChaincodeID, common.ChaincodeIDFlag, "", "The chaincode ID")
*/

type ChaincodeInfoArgs struct {
	ChannelID   string `json:"channelId"`
	ChaincodeID string `json:"chaincodeId"`
}

type getInfoAction struct {
	common.ActionImpl
}

func NewChaincodeInfoAction(args *ChaincodeInfoArgs) (*getInfoAction, error) {
	action, flags := &getInfoAction{}, &pflag.FlagSet{}

	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, args.ChannelID, "The name of channel")
	flags.StringVar(&common.ChaincodeID, common.ChaincodeIDFlag, args.ChaincodeID, "The name of chaincode")

	err := action.Initialize(flags)
	if len(action.Peers()) == 0 {
		return nil, fmt.Errorf("a peer must be specified")
	}
	return action, err
}

func (action *getInfoAction) Execute() (string, error) {
	chain, err := action.NewChain()
	if err != nil {
		return "", fmt.Errorf("Error initializing chain: %v", err)
	}

	var args []string
	args = append(args, "getccdata")
	args = append(args, common.ChannelID)
	args = append(args, common.ChaincodeID)

	cdbytes, err := common.QueryChaincode(chain, action.Peers(), "lccc", common.ChannelID, args)
	if err != nil {
		return "", fmt.Errorf("Error instantiating chaincode: %v", err)
	}

	ccData := &ccprovider.ChaincodeData{}
	err = proto.Unmarshal(cdbytes, ccData)
	if err != nil {
		return "", fmt.Errorf("Error unmarshalling chaincode data: %v", err)
	}

	action.Printer().PrintChaincodeData(ccData)

	return ccData.String(), nil
}
