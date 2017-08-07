package assetApp

import "testing"

func TestInitialize(t *testing.T) {
	if err := InitChannel(); err != nil {
		t.Fatal(err)
	}
	if err := InitCC("model_cc", "github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/fixtures/src/github.com/model_cc", "v0", []string{"Init", "init"}); err != nil {
		t.Fatal(err)
	}
}
