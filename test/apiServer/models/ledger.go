package models

import (
	"fmt"

	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
)

var ledger *sdkIgn.Ledger

type Transaction struct {
	TxId      string   `json:"txId"`
	Nonce     string   `json:"nonce"`
	Creator   string   `json:"creator"`
	Signature string   `json:"signature"`
	Endorsers []string `json:"endorsers"`
	Detail    string   `json:"detail"`
	//Payload   string   `json:"payload"`
}
type Block struct {
	Number       int    `json:"number"`
	PreviousHash string `json:"previousHash"`
	CurrentHash  string `json:"currentHash"`
	DataHash     string `json:"dataHash"`
	Data         string `json:"data"`
	Metadata     string `json:"metadata"`
}

func GetTx(txId string) (trans *Transaction, err error) {
	txInfo, err := ledger.QueryTrans(txId)
	fmt.Println(txInfo)
	if err != nil {
		return nil, err
	}
	trans = &Transaction{
		TxId:      txInfo.TxId,
		Nonce:     txInfo.Nonce,
		Creator:   txInfo.Creator,
		Signature: txInfo.Signature,
		Endorsers: txInfo.Endorsers,
		Detail:    txInfo.Detail,
		//Payload:   txInfo.Payload,
	}
	return trans, nil
}

func GetBlockByNumber(i int) (*Block, error) {
	block, err := ledger.QueryBlockByNumber(i)
	if err != nil {
		return nil, err
	}
	return &Block{
		Number:       block.Number,
		PreviousHash: block.PreHash,
		CurrentHash:  block.CurHash,
		DataHash:     block.DataHash,
		Data:         block.Data,
		Metadata:     block.Metadata,
	}, nil
}

func GetBlocks() ([]*Block, error) {
	blockchain, err := ledger.QueryBlockChain()
	if err != nil {
		return nil, err
	}
	var blocks []*Block
	for i := int(blockchain.Height - 1); i > 0; i-- {
		b, err := GetBlockByNumber(i)
		if err != nil {
			return nil, fmt.Errorf("QueryBlock [%d] err [%s]", i, err)
		}
		blocks = append(blocks, b)
	}

	gesisBlock, err := ledger.QueryBlockByNumber(0)
	if err != nil {
		return nil, fmt.Errorf("Query Genesis Block [%d] err [%s]", 0, err)
	}

	return append(blocks, &Block{
		Number:       gesisBlock.Number,
		PreviousHash: gesisBlock.PreHash,
		CurrentHash:  gesisBlock.CurHash,
		DataHash:     gesisBlock.DataHash,
		Data:         gesisBlock.Data,
		Metadata:     gesisBlock.Metadata,
	}), nil
}
