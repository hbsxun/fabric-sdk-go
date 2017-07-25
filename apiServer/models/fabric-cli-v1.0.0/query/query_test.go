package query

import "testing"

func TestQuery(t *testing.T) {
	channelID := "mychannel"
	peerUrl := "localhost:7051"

	queryBlock(&QueryBlockArgs{
		ChannelID: channelID,
		BlockNum:  "0",
		PeerUrl:   peerUrl,
	}, t)

	queryChannels(&QueryChannelsArgs{
		PeerUrl: peerUrl,
	}, t)

	queryChainInfo(&QueryChainInfoArgs{
		ChannelID: channelID,
		PeerUrl:   peerUrl,
	}, t)

	queryTx(&QueryTxArgs{
		ChannelID: channelID,
		PeerUrl:   peerUrl,
		TxID:      "", //if not set, then query all txs in 'mychannel' channel
	}, t)
}

func queryBlock(args *QueryBlockArgs, t *testing.T) {
	action, err := NewQueryBlockAction(args)
	if err != nil {
		t.Fatal(err)
	}

	if err = action.Execute(); err != nil {
		t.Fatal(err)
	}
}

func queryChannels(args *QueryChannelsArgs, t *testing.T) {
	action, err := NewQueryChannelsAction(args)
	if err != nil {
		t.Fatal(err)
	}
	if err = action.Execute(); err != nil {
		t.Fatal(err)
	}
}

func queryChainInfo(args *QueryChainInfoArgs, t *testing.T) {
	action, err := NewQueryChainInfoAction(args)
	if err != nil {
		t.Fatal(err)
	}
	if err = action.Execute(); err != nil {
		t.Fatal(err)
	}
}
func queryTx(args *QueryTxArgs, t *testing.T) {
	action, err := NewQueryTXAction(args)
	if err != nil {
		t.Fatal(err)
	}
	if err = action.Execute(); err != nil {
		t.Fatal(err)
	}
}
