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

const (
	txIDFlag = "txid"
)

var txID string

type QueryTxArgs struct {
	ChannelID string `json:"channelId"`
	txID      string `json:"txId"`
}

type queryTXAction struct {
	common.ActionImpl
}

func NewQueryTXAction(args *QueryTxArgs) (*queryTXAction, error) {
	action, flags := &queryTXAction{}, &pflag.FlagSet{}
	if args.ChannelID == "" {
		args.ChannelID = common.ChannelID
		common.Logger.Infof("using default ChannelID: %s", common.ChannelID)
	}
	flags.StringVar(&common.ChannelID, common.ChaincodeIDFlag, args.ChannelID, "The Channel ID")
	flags.StringVar(&txID, txIDFlag, args.txID, "The tx id")
	err := action.Initialize(flags)
	return action, err
}

func (action *queryTXAction) Execute() (string, error) {
	chain, err := action.NewChain()
	if err != nil {
		return "", fmt.Errorf("Error initializing chain: %v", err)
	}

	tx, err := chain.QueryTransaction(txID)
	if err != nil {
		return "", err
	}

	fmt.Printf("Transaction #%s in chain %s\n", txID, common.ChannelID)
	action.Printer().PrintProcessedTransaction(tx)

	return tx.String(), nil
}
