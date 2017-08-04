package assetApp

import "testing"

func TestInitialize(t *testing.T) {
	if err := InitChannel(); err != nil {
		t.Fatal(err)
	}
	if err := InitCC("model_cc", "v0", nil); err != nil {
		t.Fatal(err)
	}
}
