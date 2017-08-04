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

package channel

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/hyperledger/fabric-sdk-go/fabric-client/util"
	"github.com/spf13/pflag"
)

const (
	defautPeerUrl = "localhost:7051"
)

/*
	flags := chainJoinCmd.Flags()
	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, common.ChannelID, "The channel ID")
	flags.StringVar(&common.OrdererURL, common.OrdererFlag, defaultOrderer, "The URL of the orderer, e.g. localhost:7050")
	flags.String(common.PeerFlag, "", "The URL of the peer, e.g. localhost:7051")
*/

type ChannelJoinArgs struct {
	ChannelID  string `json:"channelId"`
	OrdererUrl string `json:"ordererUrl"`
	PeerUrl    string `json:"peerUrl"`
}
type channelJoinAction struct {
	common.ActionImpl
}

func NewChannelJoinAction(args *ChannelJoinArgs) (*channelJoinAction, error) {
	action, flags := &channelJoinAction{}, &pflag.FlagSet{}

	if args.ChannelID == "" {
		args.ChannelID = common.ChannelID
		common.Logger.Infof("Using default ChannelID: %s", common.ChannelID)
	}
	/*
		if args.PeerUrl == "" {
			args.PeerUrl = defautPeerUrl
			common.Logger.Infof("Using default PeerUrl: %s", defautPeerUrl)
		}
	*/
	if args.OrdererUrl == "" {
		args.OrdererUrl = defaultOrderer
		common.Logger.Infof("Using default OrdererUrl: %s", defaultOrderer)
	}
	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, args.ChannelID, "The channel ID")
	flags.StringVar(&common.PeerURL, common.PeerFlag, args.PeerUrl, "The URL of the peer, e.g. localhost:7051")
	flags.StringVar(&common.OrdererURL, common.OrdererFlag, args.OrdererUrl, "The URL of the orderer, e.g. localhost:7050")

	err := action.Initialize(flags)
	if err != nil {
		return nil, err
	}

	if len(action.Peers()) == 0 {
		return nil, fmt.Errorf("at least one peer is required for join")
	}
	return action, nil
}

func (action *channelJoinAction) Execute() error {
	chain, err := action.NewChain()
	if err != nil {
		return err
	}

	fmt.Printf("Attempting to join channel: %s\n", common.ChannelID)

	err = joinChannel(chain, action.Client(), action.Peers())
	if err != nil {
		return fmt.Errorf("Could not join channel: %v", err)
	}

	if chain != nil {
		fmt.Println("Channel joined!")
	}

	return nil
}

func joinChannel(chain fabricClient.Chain, client fabricClient.Client, peers []fabricClient.Peer) error {
	signingIdentity, err := common.GetSigningIdentity(client)
	if err != nil {
		return fmt.Errorf("Could not generate creator ID: %v", err)
	}
	nonce, err := util.GenerateRandomNonce()
	if err != nil {
		return fmt.Errorf("Could not compute nonce: %s", err)
	}
	txID, err := util.ComputeTxID(nonce, signingIdentity)
	if err != nil {
		return fmt.Errorf("Could not compute TxID: %s", err)
	}

	req := &fabricClient.JoinChannelRequest{
		Targets: peers,
		TxID:    txID,
		Nonce:   nonce,
	}

	if err = chain.JoinChannel(req); err != nil {
		return fmt.Errorf("Could not join channel: %v", err)
	}

	return nil
}
