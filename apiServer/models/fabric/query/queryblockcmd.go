/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package query

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/api/apifabclient"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/common"
	fabricCommon "github.com/hyperledger/fabric/protos/common"
	"github.com/spf13/pflag"
)

/*
var queryBlockCmd = &cobra.Command{
	Use:   "block",
	Short: "Query block",
	Long:  "Queries a block",
	Run: func(cmd *cobra.Command, args []string) {
		if common.Config().BlockNum() < 0 && common.Config().BlockHash() == "" {
			fmt.Printf("\nMust specify either the block number or the block hash\n\n")
			cmd.HelpFunc()(cmd, args)
			return
		}
		action, err := newQueryBlockAction(cmd.Flags())
		if err != nil {
			common.Config().Logger().Criticalf("Error while initializing queryBlockAction: %v", err)
			return
		}

		defer action.Terminate()

		err = action.invoke()
		if err != nil {
			common.Config().Logger().Criticalf("Error while running queryBlockAction: %v", err)
		}
	},
}

func getQueryBlockCmd() *cobra.Command {
	flags := queryBlockCmd.Flags()
	common.Config().InitChannelID(flags)
	common.Config().InitBlockNum(flags)
	common.Config().InitBlockHash(flags)
	common.Config().InitTraverse(flags)
	common.Config().InitPeerURL(flags, "", "The URL of the peer on which to install the chaincode, e.g. localhost:7051")
	return queryBlockCmd
}
*/
var blocksSlice []*fabricCommon.Block

type QueryBlockArgs struct {
	ChannelID string `json:"channelId"`
	BlockNum  string `json:"blockNum"`
	BlockHash string `json:"blockHash"`
	Traverse  string `json:"traverse"`
	PeerUrl   string `json:"peerUrl"`
}

type queryBlockAction struct {
	common.Action
}

func NewQueryBlockAction(args *QueryBlockArgs) (*queryBlockAction, error) {
	flags := &pflag.FlagSet{}
	common.Config().InitChannelID(flags, args.ChannelID)
	common.Config().InitBlockNum(flags, args.BlockNum)
	common.Config().InitBlockHash(flags, args.BlockHash)
	common.Config().InitTraverse(flags, args.Traverse)
	common.Config().InitPeerURL(flags, args.PeerUrl, "The URL of the peer on which to install the chaincode, e.g. localhost:7051")
	blocksSlice = make([]*fabricCommon.Block, 1)

	action := &queryBlockAction{}
	err := action.Initialize(flags)
	return action, err
}

func (action *queryBlockAction) Execute() ([]*fabricCommon.Block, error) {
	channelClient, err := action.ChannelClient()
	if err != nil {
		return nil, fmt.Errorf("Error getting channel client: %v", err)
	}

	context := action.SetUserContext(action.OrgAdminUser(common.Config().OrgID()))
	defer context.Restore()

	var block *fabricCommon.Block
	if common.Config().BlockNum() >= 0 {
		var err error
		block, err = channelClient.QueryBlock(common.Config().BlockNum())
		if err != nil {
			return nil, err
		}
	} else if common.Config().BlockHash() != "" {
		var err error

		hashBytes, err := common.Base64URLDecode(common.Config().BlockHash())
		if err != nil {
			return nil, err
		}

		block, err = channelClient.QueryBlockByHash(hashBytes)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("must specify either a block number of a block hash")
	}

	action.Printer().PrintBlock(block)
	blocksSlice = append(blocksSlice, block)
	action.traverse(channelClient, block, common.Config().Traverse()-1)

	return blocksSlice, nil
}

func (action *queryBlockAction) traverse(chain apifabclient.Channel, currentBlock *fabricCommon.Block, num int) error {
	if num <= 0 {
		return nil
	}

	block, err := chain.QueryBlockByHash(currentBlock.Header.PreviousHash)
	if err != nil {
		return err
	}

	action.Printer().PrintBlock(block)
	blocksSlice = append(blocksSlice, block)
	if block.Header.PreviousHash != nil {
		return action.traverse(chain, block, num-1)
	}
	return nil
}
