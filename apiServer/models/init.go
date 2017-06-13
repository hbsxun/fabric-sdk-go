package models

import (
	"os"

	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("RestServer Models")

var setup *sdkIgn.BaseSetupImpl
var prefix = os.Getenv("GOPATH") + "/src/github.com/hyperledger/fabric-sdk-go/test"

func init() {
	setup = sdkIgn.NewBaseSetupImpl(prefix)
	err := setup.InstallAndInstantiateModelCC()
	if err != nil {
		logger.Errorf("InstallAndInstantiateModelCC failed ", err)
		os.Exit(-1)
	}

	ledger = sdkIgn.NewLedger(setup.Chain)
	admin = sdkIgn.NewMember(setup.Client)
}
