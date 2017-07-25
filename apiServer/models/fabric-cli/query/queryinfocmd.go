/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package query

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/common"
	"github.com/spf13/pflag"
)

/*
var queryInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Query info",
	Long:  "Queries general info",
	Run: func(cmd *cobra.Command, args []string) {
		action, err := newQueryInfoAction(cmd.Flags())
		if err != nil {
			common.Config().Logger().Criticalf("Error while initializing queryInfoAction: %v", err)
			return
		}

		defer action.Terminate()

		err = action.run()
		if err != nil {
			common.Config().Logger().Criticalf("Error while running queryInfoAction: %v", err)
		}
	},
}

func getQueryInfoCmd() *cobra.Command {
	flags := queryInfoCmd.Flags()
	common.Config().InitTxID(flags)
	common.Config().InitChannelID(flags)
	common.Config().InitPeerURL(flags)
	return queryInfoCmd
}

type queryInfoAction struct {
	common.Action
}

func newQueryInfoAction(flags *pflag.FlagSet) (*queryInfoAction, error) {
	action := &queryInfoAction{}
	err := action.Initialize(flags)
	return action, err
}
*/
type QueryChainInfoArgs struct {
	//	TxID      string `json:"txID"`
	ChannelID string `json:"channelId"`
	PeerUrl   string `json:"peerUrl"`
}

type queryInfoAction struct {
	common.Action
}

func NewQueryChainInfoAction(args *QueryChainInfoArgs) (*queryInfoAction, error) {
	flags := &pflag.FlagSet{}
	//common.Config().InitTxID(flags, args.TxID)
	common.Config().InitChannelID(flags, args.ChannelID)
	common.Config().InitPeerURL(flags, args.PeerUrl)

	action := &queryInfoAction{}
	err := action.Initialize(flags)
	return action, err
}
func (action *queryInfoAction) Execute() error {
	channelClient, err := action.ChannelClient()
	if err != nil {
		return fmt.Errorf("Error getting channel client: %v", err)
	}

	context := action.SetUserContext(action.OrgAdminUser(action.OrgID()))
	defer context.Restore()

	info, err := channelClient.QueryInfo()
	if err != nil {
		return err
	}

	action.Printer().PrintBlockchainInfo(info)

	return nil
}
