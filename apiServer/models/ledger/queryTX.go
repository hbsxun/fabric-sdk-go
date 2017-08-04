package ledger

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric/query"
	"github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type TxInfo struct {
	TxId      string   `json:"txId"`
	Signature string   `json:"signature"`
	Creator   string   `json:"creator"`
	Nonce     string   `json:"nonce"`
	Endorsers []string `json:"endorsers"`
	Detail    string   `json:"detail"`
}

/*type TransactionInfo struct {
	Signature string      `json:"signature"`
	Payload   PayloadInfo `json:"payload"`
}
type PayloadInfo struct {
	Header HeaderInfo  `json:"header"`
	Data   interface{} `json:"data"`
}
type HeaderInfo struct {
	ChannelHeader   ChannelHeaderInfo   `json:"channelHeader"`
	SignatureHeader SignatureHeaderInfo `json:"signatureHeader"`
}
type ChannelHeaderInfo struct {
	Type      int32                      `json:"type"`
	Version   int32                      `json:"version"`
	Timestamp *google_protobuf.Timestamp `json:"timestamp"`
	ChannelId string                     `json:"channelId"`
	TxId      string                     `json:"txId"`
	Epoch     uint64                     `json:"epoch"`
	Extension ExtensionInfo              `json:"extension"`
}
type SignatureHeaderInfo struct {
	Creator string `json:"creator"`
	Nonce   string `json:"nonce"`
}
type ExtensionInfo struct {
	PayloadVisibility []byte          `json:"payloadVisibility"`
	ChaincodeId       ChaincodeIdInfo `json:"chaincodeId"`
}
type ChaincodeIdInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Path    string `json:"path"`
}
type ConfigEnvelopeInfo struct {
	Type       fabricCommon.HeaderType `json:"type"`
	Config     ConfigInfo              `json:"config"`
	LastUpdate LastUpdateInfo          `json:"lastUpdate"`
}

type ConfigInfo struct {
	Sequence     uint64           `json:"sequence"`
	ChannelGroup ChannelGroupInfo `json:"channelGroup"`
}

type ChannelGroupInfo struct {
	Version   uint64             `json:"version"`
	ModPolicy string             `json:"modPolicy"`
	Groups    []ChannelGroupInfo `json:"groups"`
	Values    []ConfigValueInfo  `json:"values"`
	Policies  []ConfigPolicyInfo `json:"policies"`
}

type ConfigValueInfo struct {
	Version   uint64      `json:"version"`
	ModPolicy string      `json:"modPolicy"`
	Detail    interface{} `json:"detail"`
}

type ConfigPolicyInfo struct {
	ModPolicy string      `json:"modPolicy"`
	Version   string      `json:"version"`
	Detail    interface{} `json:"detail"`
}*/

func QueryTX(queryTxArgs *query.QueryTxArgs) (*TxInfo, error) {
	queryTxAction, err := query.NewQueryTXAction(queryTxArgs)
	if err != nil {
		return nil, fmt.Errorf("NewQueryTXAction err [%s]", err)
	}
	processedTransaction, err := queryTxAction.Execute()
	if err != nil {
		return nil, fmt.Errorf("queryTxAction err [%s]", err)
	}
	/*transactionInfo, err := parseProcessedTransaction(processedTransaction)
	if err != nil {
		return nil, fmt.Errorf("parseProcessedTransaction err [%s]", err)
	}
	return transactionInfo, nil*/
	if processedTransaction.TransactionEnvelope == nil {
		return nil, errors.New("QueryTransaction failed to return transaction envelope")
	}
	signedData, err := processedTransaction.GetTransactionEnvelope().AsSignedData()
	if err != nil {
		return nil, fmt.Errorf("envolope.AsSignedData error [%s]", err)
	}
	var payload common.Payload
	err = proto.Unmarshal(signedData[0].Data, &payload)
	if err != nil {
		return nil, errors.New("processedTransaction Unmarshal failed, can't get payload")
	}
	creator, nonce, spec, endorsers, err := parseBody(&payload)
	if err != nil {
		return nil, err
	}
	_ = creator
	return &TxInfo{
		TxId:      queryTxArgs.TxID,
		Signature: base64.StdEncoding.EncodeToString(signedData[0].Signature),
		Creator:   base64.StdEncoding.EncodeToString(signedData[0].Identity),
		Nonce:     new(big.Int).SetBytes(nonce).String(),
		Detail:    spec,
		Endorsers: endorsers,
	}, nil
}

func parseBody(payload *common.Payload) (creator, nonce []byte, spec string, endorsers []string, err error) {
	//body
	var trans pb.Transaction
	err = proto.Unmarshal(payload.Data, &trans)
	if err != nil {
		return
	}
	/*
		for _, action := range trans.GetActions() {
			parseTxActionHeader(action.Header)
			parseTxActionPayload(action.Payload)
		}
	*/
	creator, nonce, err = parseTxActionHeader(trans.GetActions()[0].Header)
	spec, endorsers, err = parseTxActionPayload(trans.GetActions()[0].Payload)
	return
}

//parseActionHeader parse the transaction Header
func parseTxActionHeader(buf []byte) (creator, nonce []byte, err error) {
	var sigHeader common.SignatureHeader
	err = proto.Unmarshal(buf, &sigHeader)
	if err != nil {
		return nil, nil, err
	}

	return sigHeader.Creator, sigHeader.Nonce, nil
}

//parseActionPayload parse the transaction body
func parseTxActionPayload(buf []byte) (string, []string, error) {
	var ccActionPl pb.ChaincodeActionPayload
	err := proto.Unmarshal(buf, &ccActionPl)

	spec, err := parseProposal(ccActionPl.ChaincodeProposalPayload)
	if err != nil {
		return "", nil, err
	}
	endorsers, err := parseEndorsedAction(ccActionPl.GetAction())
	if err != nil {
		return "", nil, err
	}
	return spec, endorsers, nil
}

//parseProposal end
func parseProposal(buf []byte) (string, error) {
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

	return spec.String(), nil
}

//parseEndorsedAction end
func parseEndorsedAction(action *pb.ChaincodeEndorsedAction) (endorsers []string, err error) {
	//Endorsement
	for _, endorsement := range action.GetEndorsements() {
		endorsers = append(endorsers, base64.StdEncoding.EncodeToString(endorsement.Endorser))
	}
	//Proposal Response
	var responsePayload pb.ProposalResponsePayload
	var ccAction pb.ChaincodeAction
	proto.Unmarshal(action.ProposalResponsePayload, &responsePayload)
	proto.Unmarshal(responsePayload.Extension, &ccAction)

	if len(endorsers) <= 0 {
		return endorsers, errors.New("No endorserment on this transaction")
	}
	return endorsers, nil
}

/*func parseProcessedTransaction(tx *pb.ProcessedTransaction) (*TransactionInfo, error) {
	var err error
	var transactionInfo TransactionInfo
	envelope := tx.TransactionEnvelope
	transactionInfo.Signature = base64.StdEncoding.EncodeToString(envelope.Signature)
	payload := fabricUtils.ExtractPayloadOrPanic(envelope)
	transactionInfo.Payload, err = parsePayload(payload)
	if err != nil {
		return nil, fmt.Errorf("parsePayload error [%s]", err)
	}
	return &transactionInfo, nil
}

func parsePayload(payload *fabricCommon.Payload) (PayloadInfo, error) {
	var payloadInfo PayloadInfo
	chdr, err := fabricUtils.UnmarshalChannelHeader(payload.Header.ChannelHeader)
	if err != nil {
		return PayloadInfo{}, fmt.Errorf("UnmarshalChannelHeader error [%s]", err)
	}
	payloadInfo.Header.ChannelHeader = parseChannelHeader(chdr)
	sigHeader, err := fabricUtils.GetSignatureHeader(payload.Header.SignatureHeader)
	if err != nil {
		return PayloadInfo{}, fmt.Errorf("GetSignatureHeader error [%s]", err)
	}
	payloadInfo.Header.SignatureHeader = parseSignatureHeader(sigHeader)
	headerType := fabricCommon.HeaderType(chdr.Type)
	if headerType == fabricCommon.HeaderType_CONFIG {
		envelope := &fabricCommon.ConfigEnvelope{}
		if err := proto.Unmarshal(payload.Data, envelope); err != nil {
			return PayloadInfo{}, fmt.Errorf("Unmarshal ConfigEnvelope error [%s]", err)
		}
		payloadInfo.Data = parseConfigEnvelope(envelope)
	} else if headerType == fabricCommon.HeaderType_CONFIG_UPDATE {
		envelope := &fabricCommon.ConfigUpdateEnvelope{}
		if err := proto.Unmarshal(payload.Data, envelope); err != nil {
			return PayloadInfo{}, fmt.Errorf("Unmarshal ConfigUpdateEnvelope error [%s]", err)
		}
		payloadInfo.Data = parseConfigUpdateEnvelope(envelope)
	} else if headerType == fabricCommon.HeaderType_ENDORSER_TRANSACTION {
		tx, err := fabricUtils.GetTransaction(data)
		if err != nil {
			return PayloadInfo{}, fmt.Errorf("GetTransaction error [%s]", err)
		}
		payloadInfo.Data = parseTransaction(tx)
	} else {
		payloadInfo.Data = base64.StdEncoding.EncodeToString(payload.Data)
	}
	return payloadInfo, nil
}

func parseChannelHeader(chdr *fabricCommon.ChannelHeader) ChannelHeaderInfo {
	var channelHeader ChannelHeaderInfo
	channelHeader.Type = chdr.Type
	channelHeader.ChannelId = chdr.ChannelId
	channelHeader.Epoch = chdr.Epoch
	channelHeader.Timestamp = chdr.Timestamp
	channelHeader.TxId = chdr.TxId
	channelHeader.Version = chdr.Version
	ccHdrExt := &pb.ChaincodeHeaderExtension{}
	unmarshalOrPanic(chdr.Extension, ccHdrExt)
	channelHeader.Extension = parseExtension(ccHdrExt)
	return channelHeader

}

func parseExtension(ccHdrExt *pb.ChaincodeHeaderExtension) ExtensionInfo {
	var extension ExtensionInfo
	extension.PayloadVisibility = ccHdrExt.PayloadVisibility
	if ccHdrExt.ChaincodeId != nil {
		extension.ChaincodeId.Name = ccHdrExt.ChaincodeId.Name
		extension.ChaincodeId.Path = ccHdrExt.ChaincodeId.Path
		extension.ChaincodeId.Version = ccHdrExt.ChaincodeId.Version
	}
	return extension
}

func parseSignatureHeader(sigHdr *fabricCommon.SignatureHeader) SignatureHeaderInfo {
	var signatureHeader SignatureHeaderInfo
	signatureHeader.Nonce = base64.StdEncoding.EncodeToString(sigHdr.Nonce)
	signatureHeader.Creator = base64.StdEncoding.EncodeToString(sigHdr.Creator)
	return signatureHeader
}
func parseConfigEnvelope(envelope *fabricCommon.ConfigEnvelope) ConfigEnvelopeInfo {
	var configEnvelope ConfigEnvelopeInfo
	CONFIG
}
func unmarshalOrPanic(buf []byte, pb proto.Message) {
	err := proto.Unmarshal(buf, pb)
	if err != nil {
		fmt.Println("unmarshalOrPanic error")
	}
}*/
