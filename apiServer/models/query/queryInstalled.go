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

type queryInstalledAction struct {
	common.ActionImpl
}

func NewQueryInstalledAction(peerUrl string) (*queryInstalledAction, error) {
	if peerUrl == "" {
		peerUrl = "localhost:7051"
	}
	action, flags := &queryInstalledAction{}, &pflag.FlagSet{}

	flags.StringVar(&common.PeerURL, common.PeerFlag, peerUrl, "")

	err := action.Initialize(flags)
	return action, err
}

func (action *queryInstalledAction) Execute() (string, error) {
	peer := action.PeerFromURL(common.PeerURL)
	if peer == nil {
		return "", fmt.Errorf("unknown peer URL: %s", common.PeerURL)
	}

	response, err := action.Client().QueryInstalledChaincodes(peer)
	if err != nil {
		return "", err
	}

	fmt.Printf("Chaincodes for peer [%s]\n", common.PeerURL)
	action.Printer().PrintChaincodes(response.Chaincodes)
	return response.String(), nil
}
