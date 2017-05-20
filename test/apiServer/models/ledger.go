package models

import (
	"fmt"

	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
)

var ledger *sdkIgn.Ledger

func GetTx(txId string) (a *sdkIgn.TxInfo, err error) {
	txInfo, err := ledger.QueryTrans(txId)
	fmt.Println(txInfo)
	if err != nil {
		return nil, err
	}
	return txInfo, nil
}

func GetBlock() {
	//ledger.QueryBlock()
}
