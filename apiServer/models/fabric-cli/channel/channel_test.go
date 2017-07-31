package channel

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	//channelID := "mychannel"
	//ordererID := "localhost:7050"
	//txFile := "../fixtures/channel/mychannel.tx"
	//if params above aren't set, will use the default values from config_test.yaml and common/config.go
	//createAction, err := NewChannelCreateAction(&ChannelCreateArgs{TxFile: txFile})
	createAction, err := NewChannelCreateAction(&ChannelCreateArgs{
	//ChannelID: channelID,
	//TxFile:    txFile,
	//OrdererUrl: ordererID,
	})
	if err != nil {
		t.Fatal("NewChannelCreateAction failed, %v", err)
	}

	err = createAction.Execute()
	if err != nil {
		t.Fatal("CreateAction failed, %v", err)
	}

	time.Sleep(time.Second * 3)

	joinAction, err := NewChannelJoinAction(&ChannelJoinArgs{})
	if err != nil {
		t.Fatal("NewChannelJoinAction failed, %v", err)
	}

	err = joinAction.Execute()
	if err != nil {
		t.Fatal("joinAction failed, %v", err)
	}
}
