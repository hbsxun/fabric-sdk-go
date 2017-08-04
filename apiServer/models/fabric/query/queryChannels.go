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
	flags.StringVar(&peerURL, common.PeerFlag, "", "The URL of the peer to query, e.g. localhost:7051")
*/

type queryChannelsAction struct {
	common.ActionImpl
}

func NewQueryChannelsAction(args string) (*queryChannelsAction, error) {
	if args == "" {
		args = "localhost:7051"
	}
	action, flags := &queryChannelsAction{}, &pflag.FlagSet{}

	flags.StringVar(&common.PeerURL, common.PeerFlag, args, "The URL of the peer to query, e.g. localhost:7051")

	err := action.Initialize(flags)
	return action, err
}

func (action *queryChannelsAction) Execute() (string, error) {
	peer := action.PeerFromURL(common.PeerURL)
	if peer == nil {
		return "", fmt.Errorf("unknown peer URL: %s", common.PeerURL)
	}

	response, err := action.Client().QueryChannels(peer)
	if err != nil {
		return "", err
	}

	fmt.Printf("Channels for peer [%s]\n", common.PeerURL)

	action.Printer().PrintChannels(response.Channels)

	return response.String(), nil
}
