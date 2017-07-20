package event

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestListenBlock(t *testing.T) {
	action, err := NewListenBlockAction(&ListenBlockArgs{
		PeerUrl: "localhost:7051",
	})
	if err != nil {
		t.Fatal(err)
	}

	if err = action.Execute(); err != nil {
		t.Fatal(err)
	}
}

func TestListenCC(t *testing.T) {
	action, err := NewListenCCAction(&ListenCCArgs{
		PeerUrl:        "localhost:7051",
		ChaincodeID:    "example_cc",
		ChaincodeEvent: "",
	})
	if err != nil {
		t.Fatal(err)
	}

	if err = action.Execute(); err != nil {
		t.Fatal(err)
	}
}

func TestListenTx(t *testing.T) {
	action, err := NewListenTxAction(&ListenTxArgs{
		PeerUrl: "localhost:7051",
		TxID:    "",
	})
	if err != nil {
		t.Fatal(err)
	}

	if err = action.Execute(); err != nil {
		t.Fatal(err)
	}
}
