/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package query

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/spf13/pflag"
)

/*
var queryTXCmd = &cobra.Command{
	Use:   "tx",
	Short: "Query transaction",
	Long:  "Queries a transaction",
	Run: func(cmd *cobra.Command, args []string) {
		if common.Config().TxID() == "" {
			fmt.Printf("\nMust specify the transaction ID\n\n")
			cmd.HelpFunc()(cmd, args)
			return
		}
		action, err := newQueryTXAction(cmd.Flags())
		if err != nil {
			common.Config().Logger().Criticalf("Error while initializing queryTXAction: %v", err)
			return
		}

		defer action.Terminate()

		err = action.run()
		if err != nil {
			common.Config().Logger().Criticalf("Error while running queryTXAction: %v", err)
		}
	},
}

func getQueryTXCmd() *cobra.Command {
	flags := queryTXCmd.Flags()
	common.Config().InitChannelID(flags)
	common.Config().InitTxID(flags)
	common.Config().InitPeerURL(flags)
	return queryTXCmd
}
*/
type QueryTxArgs struct {
	ChannelID string `json:"channelId"`
	TxID      string `json:"txId"`
	PeerUrl   string `json:"peerUrl"`
}

type queryTXAction struct {
	common.Action
}

func NewQueryTXAction(args *QueryTxArgs) (*queryTXAction, error) {
	flags := &pflag.FlagSet{}
	common.Config().InitChannelID(flags, args.ChannelID)
	common.Config().InitTxID(flags, args.TxID)
	common.Config().InitPeerURL(flags, args.PeerUrl)

	action := &queryTXAction{}
	err := action.Initialize(flags)

	return action, err
}

func (action *queryTXAction) Execute() (*pb.ProcessedTransaction, error) {
	channelClient, err := action.ChannelClient()
	if err != nil {
		return nil, fmt.Errorf("Error getting channel client: %v", err)
	}

	context := action.SetUserContext(action.OrgAdminUser(common.Config().OrgID()))
	defer context.Restore()

	tx, err := channelClient.QueryTransaction(common.Config().TxID())
	if err != nil {
		return nil, err
	}

	fmt.Printf("Transaction %s in chain %s\n", common.Config().TxID(), common.Config().ChannelID())
	action.Printer().PrintProcessedTransaction(tx)

	return tx, nil
}
