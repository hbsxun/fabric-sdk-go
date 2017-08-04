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

package query

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	"github.com/spf13/pflag"
)

/*
	flags.StringVar(&txID, txIDFlag, "", "The transaction ID")
	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, common.ChannelID, "The channel ID")
	flags.String(common.PeerFlag, "", "The URL of the peer, e.g. localhost:7051")
*/

type QueryInfoArgs struct {
	ChannelID string `json:channelId`
	PeerUrl   string `json:peerUrl`
}

type queryInfoAction struct {
	common.ActionImpl
}

func NewQueryInfoAction(args *QueryInfoArgs) (*queryInfoAction, error) {
	action, flags := &queryInfoAction{}, &pflag.FlagSet{}

	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, args.ChannelID, "The channel ID")
	flags.StringVar(&common.PeerURL, common.PeerFlag, args.PeerUrl, "The channel ID")

	err := action.Initialize(flags)
	return action, err
}

func (action *queryInfoAction) Execute() (string, error) {
	chain, err := action.NewChain()
	if err != nil {
		return "", fmt.Errorf("Error initializing chain: %v", err)
	}

	info, err := chain.QueryInfo()
	if err != nil {
		return "", err
	}

	action.Printer().PrintBlockchainInfo(info)

	return info.String(), nil
}
