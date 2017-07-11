package AssetApp

import "testing"

func TestInitialize(t *testing.T) {
	if err := InitCC("model_cc", "v0", nil); err != nil {
		t.Fatal(err)
	}
}
