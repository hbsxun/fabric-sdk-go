package AssetApp

import (
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/channel"
)

func CreateAndJoinChannel() error {
	//create channel
	createAction, err := channel.NewChannelCreateAction(&channel.ChannelCreateArgs{})
	if err != nil {
		return err
	}
	if err = createAction.Execute(); err != nil {
		return err
	}

	//peers join channel
	joinAction, err := channel.NewChannelJoinAction(&channel.ChannelJoinArgs{})
	if err != nil {
		return err
	}
	if err = joinAction.Execute(); err != nil {
		return err
	}
	return nil
}
