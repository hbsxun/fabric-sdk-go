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
	"math/big"

	"github.com/golang/protobuf/proto"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("Fabric Ledger Query")

type TxInfo struct {
	TxId      string   `json:"txId"`
	Signature string   `json:"signature"`
	Creator   string   `json:"creator"`
	Nonce     string   `json:"nonce"`
	Endorsers []string `json:"endorsers"`
	Detail    string   `json:"detail"`
	Payload   string   `json:"payload"`
}

type BlockInfo struct {
	Number   int    `json:"number"`
	PreHash  string `json:"preHash"`
	CurHash  string `json:"curHash"`
	DataHash string `json:"dataHash"`
	Payload  string `json:"payload"`
}

//Ledger
type Ledger struct {
	chain fabricClient.Chain
}

//NewLedger returns an instance of Ledger
func NewLedger(chain fabricClient.Chain) *Ledger {
	return &Ledger{chain}
}

func (t *TxInfo) String() string {
	return fmt.Sprintf("Transaction ID: %s \nCreator Signature: %s \nCreator Identity: %s \nTransaction Nonce: %s \nEndorsers: %v \nTransaction Content: %s \npayload %s\n", t.TxId, t.Signature, t.Creator, t.Nonce, t.Endorsers, t.Detail, t.Payload)
}

func (t *BlockInfo) String() string {
	return fmt.Sprintf("Number: %d \nPreHash: %s \nCurHash: %s \nDataHash: %s \npayload %s\n", t.Number, t.PreHash, t.CurHash, t.DataHash, t.Payload)
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
		TxId:      txID,
		Signature: base64.StdEncoding.EncodeToString(signedData[0].Signature),
		Creator:   base64.StdEncoding.EncodeToString(signedData[0].Identity),
		Nonce:     new(big.Int).SetBytes(nonce).String(),
		Detail:    spec,
		Endorsers: endorsers,
		Payload:   processedTransaction.String(),
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

//QueryBlockChain {Height, CurrentHash, PreviousHash}
func (this *Ledger) QueryBlockChain() (*common.BlockchainInfo, error) {
	// Retrieve current blockchain info
	bci, err := this.chain.QueryInfo()
	if err != nil {
		return nil, fmt.Errorf("QueryInfo return error: %v", err)
	}
	return bci, nil
}

//QueryBlockByNumber
func (this *Ledger) QueryBlockByNumber(number int) (*BlockInfo, error) {
	block, err := this.chain.QueryBlock(number)
	if err != nil {
		return nil, fmt.Errorf("QueryBlock return error [%s]", err)
	}

	return &BlockInfo{
		Number:   int(block.GetHeader().Number),
		CurHash:  base64.StdEncoding.EncodeToString(block.GetHeader().Hash()),
		PreHash:  base64.StdEncoding.EncodeToString(block.GetHeader().PreviousHash),
		DataHash: base64.StdEncoding.EncodeToString(block.GetHeader().DataHash),
		Payload:  block.String(),
	}, nil
}

//QueryBlockByHash
func (this *Ledger) QueryBlockByHash(hash []byte) (*BlockInfo, error) {
	block, err := this.chain.QueryBlockByHash(hash)
	if err != nil {
		return nil, fmt.Errorf("QueryBlockByHash return error [%s]", err)
	}

	return &BlockInfo{
		Number:   int(block.GetHeader().Number),
		CurHash:  base64.StdEncoding.EncodeToString(block.GetHeader().Hash()),
		PreHash:  base64.StdEncoding.EncodeToString(block.GetHeader().PreviousHash),
		DataHash: base64.StdEncoding.EncodeToString(block.GetHeader().DataHash),
		Payload:  block.String(),
	}, nil
}

//QueryBlocks
func (this *Ledger) QueryBlocks() (blocks []*BlockInfo, err error) {
	blockchain, err := this.QueryBlockChain()
	if err != nil {
		return nil, err
	}
	for i := blockchain.Height - 1; i > 0; i-- {
		block, err := this.QueryBlockByNumber(int(i))
		if err != nil {
			return nil, err
		}
		blocks = append(blocks, block)
	}
	return blocks, nil
}
