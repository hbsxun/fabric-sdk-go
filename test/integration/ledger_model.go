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
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/golang/protobuf/proto"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("Fabric Ledger Query")

type TxInfo struct {
	txId           string
	signature      string
	creator, nonce string
	endorsers      []string
	detail         string
}

func (t *TxInfo) String() string {
	return fmt.Sprintf("Transaction ID: %s \nCreator Signature: %s \nCreator Identity: %s \nTransaction Nonce: %s \nEndorsers: %v \nTransaction Content: %s \n", t.txId, t.signature, t.creator, t.nonce, t.endorsers, t.detail)
}

type BlockInfo struct {
	blockNumber      int
	preHash, curHash string
	txNumber         int
	txHashs          []string
}

type Ledger struct {
	chain fabricClient.Chain
}

func NewLedger(chain fabricClient.Chain) *Ledger {
	return &Ledger{chain}
}

func (this *Ledger) QueryTrans(txID string) (tx *TxInfo, err error) {
	chain := this.chain
	// Test Query Transaction -- verify that valid transaction has been processed
	processedTransaction, err := chain.QueryTransaction(txID)
	if err != nil {
		return nil, fmt.Errorf("QueryTransaction return error: %v", err)
	}
	if processedTransaction.TransactionEnvelope == nil {
		return nil, errors.New("QueryTransaction failed to return transaction envelope")
	}

	//common Envelope{payload, Signature}
	logger.Debugf("txID [%s]\n ", txID)
	logger.Debugf("Transaction ValidationCode_name is [%s]\n", pb.TxValidationCode_name[processedTransaction.ValidationCode])
	//signedData is a slice which length is 1
	signedData, err := processedTransaction.GetTransactionEnvelope().AsSignedData()
	if err != nil {
		return nil, fmt.Errorf("envolope.AsSignedData error [%s]", err)
	}
	//logger.Debugf("signedData payload %v\n ", signedData[0].Data)
	logger.Debugf("signedData Signature %v\n ", base64.StdEncoding.EncodeToString(signedData[0].Signature))
	//logger.Debugf("signedData creator %v\n ", string(signedData[0].Identity))

	//payload
	var payload common.Payload
	err = proto.Unmarshal(signedData[0].Data, &payload)
	if err != nil {
		return nil, errors.New("processedTransaction Unmarshal failed, can't get payload")
	}

	//parse header
	//this.parseHeader(&payload)

	//parse body
	creator, nonce, spec, endorsers, err := this.parseBody(&payload)
	if err != nil {
		return nil, err
	}
	_ = creator
	return &TxInfo{
		txId:      txID,
		signature: base64.StdEncoding.EncodeToString(signedData[0].Signature),
		creator:   string(signedData[0].Identity),
		nonce:     new(big.Int).SetBytes(nonce).String(),
		detail:    spec,
		endorsers: endorsers,
	}, nil
}

//parseHeader parse the processedTransaction header
func (this *Ledger) parseHeader(payload *common.Payload) {
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
func (this *Ledger) parseBody(payload *common.Payload) (creator, nonce []byte, spec string, endorsers []string, err error) {
	//body
	var trans pb.Transaction
	err = proto.Unmarshal(payload.Data, &trans)
	if err != nil {
		return
	}
	/*
		for _, action := range trans.GetActions() {
			this.parseTxActionHeader(action.Header)
			this.parseTxActionPayload(action.Payload)
		}
	*/
	creator, nonce, err = this.parseTxActionHeader(trans.GetActions()[0].Header)
	spec, endorsers, err = this.parseTxActionPayload(trans.GetActions()[0].Payload)
	return
}

//parseActionHeader parse the transaction Header
func (this *Ledger) parseTxActionHeader(buf []byte) (creator, nonce []byte, err error) {
	var sigHeader common.SignatureHeader
	err = proto.Unmarshal(buf, &sigHeader)
	if err != nil {
		return nil, nil, err
	}

	//sigHeader: creator, nonce
	logger.Debug(sigHeader.String())

	return sigHeader.Creator, sigHeader.Nonce, nil
}

//parseActionPayload parse the transaction body
func (this *Ledger) parseTxActionPayload(buf []byte) (string, []string, error) {
	var ccActionPl pb.ChaincodeActionPayload
	err := proto.Unmarshal(buf, &ccActionPl)

	//ccActionPl: ChaincodeProposalPayload, Actions
	logger.Debug(ccActionPl.String())

	spec, err := this.parseProposal(ccActionPl.ChaincodeProposalPayload)
	if err != nil {
		return "", nil, err
	}
	endorsers, err := this.parseEndorsedAction(ccActionPl.GetAction())
	if err != nil {
		return "", nil, err
	}
	return spec, endorsers, nil
}

//parseProposal end
func (this *Ledger) parseProposal(buf []byte) (string, error) {
	var proposal pb.ChaincodeProposalPayload
	err := proto.Unmarshal(buf, &proposal)
	if err != nil {
		return "", err
	}

	//conclude: the proposal.Input is a ChaincodeXXXXSpec
	var spec pb.ChaincodeInvocationSpec
	err = proto.Unmarshal(proposal.Input, &spec)
	if err != nil {
		return "", err
	}

	logger.Debugf("Proposal Input: %v\n", spec.String())

	return spec.String(), nil
}

//parseEndorsedAction end
func (this *Ledger) parseEndorsedAction(action *pb.ChaincodeEndorsedAction) (endorsers []string, err error) {
	//Endorsement
	for _, endorsement := range action.GetEndorsements() {
		endorsers = append(endorsers, string(endorsement.Endorser))
	}
	//Proposal Response
	var responsePayload pb.ProposalResponsePayload
	var ccAction pb.ChaincodeAction
	proto.Unmarshal(action.ProposalResponsePayload, &responsePayload)
	proto.Unmarshal(responsePayload.Extension, &ccAction)

	logger.Debugf("Result %s\n", ccAction.String())

	if len(endorsers) <= 0 {
		return endorsers, errors.New("No endorserment on this transaction")
	}
	return endorsers, nil
}

//QueryBlock
func (this *Ledger) QueryBlock() {
	chain := this.chain
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
	for i := bci.Height - 1; i > 0; i-- {
		block, err := chain.QueryBlock(int(i))
		if err != nil {
			logger.Fatalf("QueryBlock return error [%s]", err)
		}
		curHash := base64.StdEncoding.EncodeToString(block.GetHeader().Hash())
		preHash := base64.StdEncoding.EncodeToString(block.GetHeader().PreviousHash)
		logger.Debugf("Height [%d] \nCurrentBlockHash [%s] \nPreviousBlockHash [%s]\n\n", i, curHash, preHash)
	}

}
func (this *Ledger) QueryGenesisBlock() {
	chain := this.chain
	i := 0
	block, err := chain.QueryBlock(i)
	if err != nil {
		logger.Fatalf("QueryBlock return error [%s]", err)
	}
	curHash := base64.StdEncoding.EncodeToString(block.GetHeader().Hash())
	preHash := base64.StdEncoding.EncodeToString(block.GetHeader().PreviousHash)
	logger.Debugf("Height [%d] \nCurrentBlockHash [%s] \nPreviousBlockHash [%s]\n\n", i, curHash, preHash)
	//logger.Debugf("%s\n\n%s\n\n%s\n\n", block.GetHeader().String(), block.GetData().String(), block.GetMetadata().String())
}
