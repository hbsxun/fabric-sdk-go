/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package query

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/spf13/pflag"
)

/*
var queryInstalledCmd = &cobra.Command{
	Use:   "installed",
	Short: "Query installed chaincodes",
	Long:  "Queries the chaincodes installed to the specified peer",
	Run: func(cmd *cobra.Command, args []string) {
		if common.Config().PeerURL() == "" {
			fmt.Printf("\nMust specify the peer URL\n\n")
			cmd.HelpFunc()(cmd, args)
			return
		}
		action, err := newqueryInstalledAction(cmd.Flags())
		if err != nil {
			common.Config().Logger().Criticalf("Error while initializing queryInstalledAction: %v", err)
			return
		}

		defer action.Terminate()

		err = action.run()
		if err != nil {
			common.Config().Logger().Criticalf("Error while running queryInstalledAction: %v", err)
		}
	},
}

func getQueryInstalledCmd() *cobra.Command {
	common.Config().InitPeerURL(queryInstalledCmd.Flags())
	return queryInstalledCmd
}
*/
type QueryInstalledArgs struct {
	PeerUrl string `json:"peerUrl"`
}

type queryInstalledAction struct {
	common.Action
}

func NewQueryInstalledAction(args *QueryInstalledArgs) (*queryInstalledAction, error) {
	flags := &pflag.FlagSet{}
	common.Config().InitPeerURL(flags, args.PeerUrl)

	action := &queryInstalledAction{}
	err := action.Initialize(flags)
	return action, err
}

func (action *queryInstalledAction) Execute() ([]*pb.ChaincodeInfo, error) {
	peer := action.PeerFromURL(common.Config().PeerURL())
	if peer == nil {
		return nil, fmt.Errorf("unknown peer URL: %s", common.Config().PeerURL())
	}

	orgID, err := action.OrgOfPeer(peer.URL())
	if err != nil {
		return nil, err
	}

	context := action.SetUserContext(action.OrgAdminUser(orgID))
	defer context.Restore()

	response, err := action.Client().QueryInstalledChaincodes(peer)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Chaincodes for peer [%s]\n", peer.URL())
	action.Printer().PrintChaincodes(response.Chaincodes)
	return response.Chaincodes, nil
}
