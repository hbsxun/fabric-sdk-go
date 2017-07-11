package AssetApp

import (
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/chaincode"
)

var chaincodeId string

type AdminCCImpl struct{}

func (this *AdminCCImpl) Deploy(args *chaincode.InstallCCArgs) error {

	chaincodeId = args.ChaincodeName

	installAction, err := chaincode.NewInstallAction(args)
	if err != nil {
		return err
	}
	if err = installAction.Execute(); err != nil {
		return err
	}
	return nil
}

func (this *AdminCCImpl) Instantiate(args *chaincode.InstantiateArgs) (string, error) {
	action, err := chaincode.NewInstantiateAction(args)
	if err != nil {
		return "", err
	}
	txId, err := action.Execute()
	if err != nil {
		return "", err
	}
	return txId, nil
}

func (this *AdminCCImpl) RegisterAsset(args *chaincode.InvokeArgs) (string, error) {

	return "", nil
}
