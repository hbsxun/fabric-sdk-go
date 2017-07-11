package AssetApp

import "testing"

func TestChannel(t *testing.T) {
	if err := InitChannel(); err != nil {
		t.Fatal(err)
	}
}
