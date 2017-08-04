package ledger

import (
	"encoding/base64"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/query"
)

type BlockInfo struct {
	Number   int    `json:"number"`
	PreHash  string `json:"preHash"`
	CurHash  string `json:"curHash"`
	DataHash string `json:"dataHash"`
	Data     string `json:"data"`
	Metadata string `json:"metadata"`
}

func QueryBlocks(queryBlocksArgs *query.QueryBlockArgs) ([]*BlockInfo, error) {
	queryBlocksAction, err := query.NewQueryBlockAction(queryBlocksArgs)
	if err != nil {
		return nil, fmt.Errorf("NewQueryBlockAction err [%s]", err)
	}
	blocks, err := queryBlocksAction.Execute()
	if err != nil {
		return nil, fmt.Errorf("queryBlocksAction err [%s]", err)
	}
	var blockInfo *BlockInfo
	var blocksInfo []*BlockInfo
	for i := 1; i < len(blocks); i++ {
		blockInfo = &BlockInfo{
			Number:   int(blocks[i].GetHeader().Number),
			CurHash:  base64.StdEncoding.EncodeToString(blocks[i].GetHeader().Hash()),
			PreHash:  base64.StdEncoding.EncodeToString(blocks[i].GetHeader().PreviousHash),
			DataHash: base64.StdEncoding.EncodeToString(blocks[i].GetHeader().DataHash),
			Data:     blocks[i].GetData().String(),
			Metadata: blocks[i].GetMetadata().String(),
		}
		blocksInfo = append(blocksInfo, blockInfo)
	}
	return blocksInfo, nil
}
