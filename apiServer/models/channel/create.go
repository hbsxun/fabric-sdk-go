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
	"io/ioutil"
	"os"

	"github.com/hyperledger/fabric-sdk-go/config"
	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/spf13/pflag"
)

const (
	txFileFlag     = "txfile"
	defaultTxFile  = "fixtures/channel/testchannel.tx"
	defaultOrderer = "localhost:7050"
)

var txFile string

type ChannelCreateArgs struct {
	ChannelID  string `json:"channelId"`
	TxFile     string `json:"txFile"`
	OrdererUrl string `json:"ordererUrl"`
}

type channelCreateAction struct {
	common.ActionImpl
}

func NewChannelCreateAction(args *ChannelCreateArgs) (*channelCreateAction, error) {
	action, flags := &channelCreateAction{}, &pflag.FlagSet{}
	//note the channelId contained in Txfile, so they should consitent
	if args.ChannelID == "" || args.TxFile == "" || args.OrdererUrl == "" {
		args.ChannelID = common.ChannelID
		common.Logger.Infof("Using default ChannelID: %s", common.ChannelID)

		args.TxFile = os.Getenv("GOPATH") + "/src/github.com/hyperledger/fabric-sdk-go/fabric-cli/" + defaultTxFile
		common.Logger.Infof("Using default TxFile: %s", defaultTxFile)

		args.OrdererUrl = defaultOrderer
		common.Logger.Infof("Using default OrdererUrl: %s", defaultOrderer)
	}
	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, args.ChannelID, "The channel ID")
	flags.StringVar(&txFile, txFileFlag, args.TxFile, "The path of the channel.tx file")
	flags.StringVar(&common.OrdererURL, common.OrdererFlag, args.OrdererUrl, "The URL of the orderer, e.g. localhost:7050")

	err := action.Initialize(flags)
	return action, err
}

func (action *channelCreateAction) Execute() error {
	configTx, err := ioutil.ReadFile(txFile)
	if err != nil {
		return fmt.Errorf("An error occurred while reading TX file %s: %v", txFile, err)
	}

	certificate := config.GetOrdererTLSCertificate()
	serverHostOverride := "orderer0"

	orderer, err := fabricClient.NewOrderer(common.OrdererURL, certificate, serverHostOverride)
	if err != nil {
		return fmt.Errorf("CreateNewOrderer return error: %v", err)
	}

	fmt.Printf("Attempting to create channel: %s\n", common.ChannelID)

	chain, err := action.Client().CreateChannel(&fabricClient.CreateChannelRequest{
		Envelope: configTx,
		Orderer:  orderer,
		Name:     common.ChannelID,
	})
	if err != nil {
		return fmt.Errorf("Error from create channel: %s", err.Error())
	}

	if chain != nil {
		fmt.Println("Channel created!")
	}

	return nil
}
