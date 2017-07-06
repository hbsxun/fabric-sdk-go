package channel

func CreateAndJoinChannel(){
	createAction, _ := NewChannelCreateAction(&ChannelCreateArgs{})
	createAction.Execute()
	joinAction, _ := NewChannelJoinAction(&ChannelJoinArgs{})
	joinAction.Execute()
}
