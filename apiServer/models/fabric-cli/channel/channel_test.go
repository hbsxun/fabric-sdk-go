package channel

import (
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	createAction, err := NewChannelCreateAction(&ChannelCreateArgs{})
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
