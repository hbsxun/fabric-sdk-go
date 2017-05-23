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
}
type Block struct {
	Number       int64  `json:"number"`
	PreviousHash string `json:"previousHash"`
	CurrentHash  string `json:"currentHash"`
	DataHash     string `json:"dataHash"`
}

var BlockList []*Blockw

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
	}
	return trans, nil
}

func GetBlocks() {
	//ledger.QueryBlock()
}
