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
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	fabricCommon "github.com/hyperledger/fabric/protos/common"
	"github.com/spf13/pflag"
)

const (
	blockNumFlag  = "num"
	blockHashFlag = "hash"
	traverseFlag  = "traverse"
)

var blockNum int
var blockHash string
var traverse int = 0
var blocksSlice []*fabricCommon.Block

/*
	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, common.ChannelID, "The channel ID")
	flags.IntVar(&blockNum, blockNumFlag, -1, "The block number")
	flags.StringVar(&blockHash, blockHashFlag, "", "The block hash")
	flags.IntVar(&traverse, traverseFlag, 0, "Blocks will be traversed starting with the given block in reverse order up to the given number of blocks")
	flags.String(common.PeerFlag, "", "The URL of the peer on which to install the chaincode, e.g. localhost:7051")
*/

type QueryBlockArgs struct {
	ChannelID string `json:"channelId"`
	BlockNum  int    `json:"blockNum"`
	BlockHash string `json:"blockHash"`
	Traverse  int    `json:"traverse"`
	PeerUrl   string `json:"peerUrl"`
}

type queryBlockAction struct {
	common.ActionImpl
}

func NewQueryBlockAction(args *QueryBlockArgs) (*queryBlockAction, error) {
	action, flags := &queryBlockAction{}, &pflag.FlagSet{}

	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, args.ChannelID, "The channel ID")
	flags.IntVar(&blockNum, blockNumFlag, args.BlockNum, "The block number")
	flags.StringVar(&blockHash, blockHashFlag, args.BlockHash, "The block hash")
	flags.IntVar(&traverse, traverseFlag, args.Traverse, "The number of blocks to traverse")
	flags.StringVar(&common.PeerURL, common.PeerFlag, args.PeerUrl, "The URL of the peer on which to query block, e.g. localhost:7051")

	blocksSlice = make([]*fabricCommon.Block, 1)
	err := action.Initialize(flags)
	return action, err
}

func (action *queryBlockAction) Execute() ([]*fabricCommon.Block, error) {
	chain, err := action.NewChain()
	if err != nil {
		return nil, fmt.Errorf("Error initializing chain: %v", err)
	}

	var block *fabricCommon.Block
	if blockNum >= 0 {
		var err error
		block, err = chain.QueryBlock(blockNum)
		if err != nil {
			return nil, err
		}
	} else if blockHash != "" {
		var err error

		hashBytes, err := common.Base64URLDecode(blockHash)
		if err != nil {
			return nil, err
		}

		block, err = chain.QueryBlockByHash(hashBytes)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("must specify either a block number or a block hash")
	}

	action.Printer().PrintBlock(block)
	blocksSlice = append(blocksSlice, block)

	action.traverse(chain, block, traverse-1)

	return blocksSlice, nil
}

func (action *queryBlockAction) traverse(chain fabricClient.Chain, currentBlock *fabricCommon.Block, num int) error {
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
