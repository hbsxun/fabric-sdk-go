package AssetApp

import "testing"

func Test_CreateAndJoinChannel(t *testing.T) {
	if err := CreateAndJoinChannel(); err != nil {
		t.Errorf("CreateAndJoinChannel err [%v]\n", err)
	}
}
