package channel

import (
	"testing"
	"time"
)

func TestCreateAndJoinChannel(t *testing.T) {
	create(t)
	time.Sleep(time.Second * 5) //must wait for a while until channel is created
	join(t)
}

func create(t *testing.T) {
	createAction, err := NewChannelCreateAction(&ChannelCreateArgs{})
	if err != nil {
		t.Fatal(err)
	}
	if err = createAction.Execute(); err != nil {
		t.Fatal(err)
	}
}

func join(t *testing.T) {
	joinAction, err := NewChannelJoinAction(&ChannelJoinArgs{})
	if err != nil {
		t.Fatal(err)
	}
	if err = joinAction.Execute(); err != nil {
		t.Fatal(err)
	}
}
