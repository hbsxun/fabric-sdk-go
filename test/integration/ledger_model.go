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

package integration

import (
	"log"

	proto "github.com/golang/protobuf/proto"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	pb "github.com/hyperledger/fabric/protos/common"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("Fabric Ledger Query")

type ledger struct {
}

func NewLedger() *ledger {
	return &ledger{}
}

func (this *ledger) QueryTrans(chain fabricClient.Chain, txID string) {

	// Test Query Transaction -- verify that valid transaction has been processed
	processedTransaction, err := chain.QueryTransaction(txID)
	if err != nil {
		log.Fatalf("QueryTransaction return error: %v", err)
	}
	//print the query transaction
	//logger.Debugf("transaction [%s]\n%v", txID, processedTransaction)

	if processedTransaction.TransactionEnvelope == nil {
		log.Fatalf("QueryTransaction failed to return transaction envelope")
	}

	envolope := processedTransaction.GetTransactionEnvelope()
	var payload pb.Payload
	err = proto.Unmarshal(envolope.Payload, &payload)
	if err != nil {
		log.Fatalf("processedTransaction Unmarshal failed, can't get payload")
	}
	//TODO to be parse further, payload: Header and Data
	logger.Debugf("txID [%s]\n Payload.Header [%v]\n Payload.Data [%s]\n Signature [%s]\n", txID, payload.GetHeader(), string(payload.Data), string(envolope.Signature))
}

func (this *ledger) QueryBlock(chain fabricClient.Chain) {

	// Retrieve current blockchain info
	bci, err := chain.QueryInfo()
	if err != nil {
		log.Fatalf("QueryInfo return error: %v", err)
	}

	// Query Block by Hash - retrieve current block by hash
	block, err := chain.QueryBlockByHash(bci.CurrentBlockHash)
	if err != nil {
		log.Fatalf("QueryBlockByHash return error: %v", err)
	}
	//print the current block
	logger.Debugf("bci [%v]\n block [%v]\n", bci, block)

	if block.Data == nil {
		log.Fatalf("QueryBlockByHash block data is nil")
	}
}
