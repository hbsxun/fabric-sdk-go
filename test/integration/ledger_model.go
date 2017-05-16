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
	"encoding/base64"
	"log"

	"github.com/golang/protobuf/proto"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
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
	if processedTransaction.TransactionEnvelope == nil {
		log.Fatalf("QueryTransaction failed to return transaction envelope")
	}

	//common Envelope{payload, Signature}
	logger.Debugf("txID [%s]\n ", txID)
	logger.Debugf("Transaction ValidationCode_name is [%s]\n", pb.TxValidationCode_name[processedTransaction.ValidationCode])
	//signedData is a slice which length is 1
	signedData, err := processedTransaction.GetTransactionEnvelope().AsSignedData()
	if err != nil {
		log.Fatalf("envolope.AsSignedData error [%s]", err)
	}
	//logger.Debugf("signedData payload %v\n ", signedData[0].Data)
	logger.Debugf("signedData Signature %v\n ", base64.StdEncoding.EncodeToString(signedData[0].Signature))
	//logger.Debugf("signedData creator %v\n ", string(signedData[0].Identity))

	//payload
	var payload common.Payload
	err = proto.Unmarshal(signedData[0].Data, &payload)
	if err != nil {
		log.Fatalf("processedTransaction Unmarshal failed, can't get payload")
	}

	//parse header
	this.parseHeader(&payload)

	//parse body
	this.parseBody(&payload)
}

//parseHeader parse the processedTransaction header
func (this *ledger) parseHeader(payload *common.Payload) {
	//header
	var header *common.Header
	var chanHeader common.ChannelHeader
	//var sigHeader common.SignatureHeader

	header = payload.GetHeader()
	_ = proto.Unmarshal(header.ChannelHeader, &chanHeader)
	//logger.Debugf("ChannelHeader %v\n ", chanHeader)

	//returns ENDORSER_TRANSACTION
	logger.Debugf("Type [%d]\n", common.HeaderType_name[chanHeader.Type])
}

//parseBody parse the processedTransaction data (that is Transaction)
func (this *ledger) parseBody(payload *common.Payload) {
	//body
	var trans pb.Transaction
	proto.Unmarshal(payload.Data, &trans)
	for _, action := range trans.GetActions() {
		this.parseTxActionHeader(action.Header)
		this.parseTxActionPayload(action.Payload)
	}
}

//parseActionHeader parse the transaction Header
func (this *ledger) parseTxActionHeader(buf []byte) {
	var sigHeader common.SignatureHeader
	proto.Unmarshal(buf, &sigHeader)

	//sigHeader: creator, nonce
	logger.Debug(sigHeader.String())
}

//parseActionPayload parse the transaction body
func (this *ledger) parseTxActionPayload(buf []byte) {
	var ccActionPl pb.ChaincodeActionPayload
	proto.Unmarshal(buf, &ccActionPl)

	//ccActionPl: ChaincodeProposalPayload, Actions
	logger.Debug(ccActionPl.String())

	//this.parseProposal(ccActionPl.ChaincodeProposalPayload)
	//this.parseEndorsedAction(ccActionPl.GetAction())
}

//parseProposal end
func (this *ledger) parseProposal(buf []byte) {
	var proposal pb.ChaincodeProposalPayload
	proto.Unmarshal(buf, &proposal)

	//conclude: the proposal.Input is a ChaincodeXXXXSpec
	var spec pb.ChaincodeInvocationSpec
	proto.Unmarshal(proposal.Input, &spec)
	logger.Debugf("Proposal Input: %v\n", spec.String())
}

//parseEndorsedAction end
func (this *ledger) parseEndorsedAction(action *pb.ChaincodeEndorsedAction) {
	//Endorsement
	for _, endorsement := range action.GetEndorsements() {
		logger.Debugf("Endorser[%s] \nSignature[%s] \n", string(endorsement.Endorser), base64.StdEncoding.EncodeToString(endorsement.Signature))
	}
	//Proposal Response
	var responsePayload pb.ProposalResponsePayload
	var ccAction pb.ChaincodeAction
	proto.Unmarshal(action.ProposalResponsePayload, &responsePayload)
	proto.Unmarshal(responsePayload.Extension, &ccAction)

	logger.Debugf("Result %s\n", ccAction.String())
}

//QueryBlock
func (this *ledger) QueryBlock(chain fabricClient.Chain) {

	// Retrieve current blockchain info
	bci, err := chain.QueryInfo()
	if err != nil {
		log.Fatalf("QueryInfo return error: %v", err)
	}
	logger.Debugf("%s\n\n", bci.String())
	/*
		// Query Block by Hash - retrieve current block by hash
		block, err := chain.QueryBlockByHash(bci.CurrentBlockHash)
		if err != nil {
			log.Fatalf("QueryBlockByHash return error: %v", err)
		}
		//print the current block
		logger.Debugf("%s\n\n%s\n\n%s\n\n", block.GetHeader().String(), block.GetData().String(), block.GetMetadata().String())
	*/
	for i := bci.Height - 1; i >= 0; i-- {
		block, err := chain.QueryBlock(int(i))
		if err != nil {
			logger.Fatalf("QueryBlock return error [%s]", err)
		}
		curHash := base64.StdEncoding.EncodeToString(block.GetHeader().Hash())
		preHash := base64.StdEncoding.EncodeToString(block.GetHeader().PreviousHash)
		logger.Debugf("Height [%d] \nCurrentBlockHash [%s] \nPreviousBlockHash [%s]\n\n", i, curHash, preHash)
	}

}
func (this *ledger) QueryGenesisBlock(chain fabricClient.Chain) {
	i := 0
	block, err := chain.QueryBlock(i)
	if err != nil {
		logger.Fatalf("QueryBlock return error [%s]", err)
	}
	curHash := base64.StdEncoding.EncodeToString(block.GetHeader().Hash())
	preHash := base64.StdEncoding.EncodeToString(block.GetHeader().PreviousHash)
	logger.Debugf("Height [%d] \nCurrentBlockHash [%s] \nPreviousBlockHash [%s]\n\n", i, curHash, preHash)
	logger.Debugf("%s\n\n%s\n\n%s\n\n", block.GetHeader().String(), block.GetData().String(), block.GetMetadata().String())
}
