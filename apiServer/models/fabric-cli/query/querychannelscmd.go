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
var queryChannelsCmd = &cobra.Command{
	Use:   "channels",
	Short: "Query channels",
	Long:  "Queries the channels of the specified peer",
	Run: func(cmd *cobra.Command, args []string) {
		if common.Config().PeerURL() == "" {
			fmt.Printf("\nMust specify the peer URL\n\n")
			cmd.HelpFunc()(cmd, args)
			return
		}
		action, err := newQueryChannelsAction(cmd.Flags())
		if err != nil {
			common.Config().Logger().Criticalf("Error while initializing queryChannelsAction: %v", err)
			return
		}

		defer action.Terminate()

		err = action.run()
		if err != nil {
			common.Config().Logger().Criticalf("Error while running queryChannelsAction: %v", err)
		}
	},
}

func getQueryChannelsCmd() *cobra.Command {
	common.Config().InitPeerURL(queryChannelsCmd.Flags())
	return queryChannelsCmd
}
*/

type QueryChannelsArgs struct {
	PeerUrl string `json:"peerUrl"`
}

type queryChannelsAction struct {
	common.Action
}

func NewQueryChannelsAction(args *QueryChannelsArgs) (*queryChannelsAction, error) {
	flags := &pflag.FlagSet{}
	common.Config().InitPeerURL(flags, args.PeerUrl)

	action := &queryChannelsAction{}
	err := action.Initialize(flags)
	return action, err
}

func (action *queryChannelsAction) Execute() error {
	peer := action.PeerFromURL(common.Config().PeerURL())
	if peer == nil {
		return fmt.Errorf("unknown peer URL: %s", common.Config().PeerURL())
	}

	orgID, err := action.OrgOfPeer(peer.URL())
	if err != nil {
		return err
	}

	context := action.SetUserContext(action.OrgAdminUser(orgID))
	defer context.Restore()

	response, err := action.Client().QueryChannels(peer)
	if err != nil {
		return err
	}

	fmt.Printf("Channels for peer [%s]\n", peer.URL())

	action.Printer().PrintChannels(response.Channels)

	return nil
}
