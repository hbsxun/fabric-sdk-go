package assetApp

import (
	"time"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/chaincode"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/channel"
)

func Initialize() error {
	err := InitChannel()
	if err != nil {
		appLogger.Debugf("CreateChannel err [%v]\n", err)
		return err
	}
	err = InitCC("model_cc", "v0", nil)
	if err != nil {
		appLogger.Debugf("InitCC err [%v]\n", err)
		return err
	}
	return nil
}

//InitChannel create and join channel
func InitChannel() error {
	//create channel
	createAction, err := channel.NewChannelCreateAction(&channel.ChannelCreateArgs{})
	if err != nil {
		appLogger.Debugf("NewChannelCreateAction err [%v]\n", err)
		return err
	}
	if err = createAction.Execute(); err != nil {
		appLogger.Debugf("createAction err [%v]\n", err)
		return err
	}

	time.Sleep(time.Second * 5)

	//peers join channel
	joinAction, err := channel.NewChannelJoinAction(&channel.ChannelJoinArgs{})
	if err != nil {
		appLogger.Debugf("NewChannelJoinAction err [%v]\n", err)
		return err
	}
	if err = joinAction.Execute(); err != nil {
		appLogger.Debugf("joinAction err [%v]\n", err)
		return err
	}
	return nil

}

//InitCC install and Initialize cc
func InitCC(chaincodeID, chaincodeVersion string, args []string) error {
	//install chaincode on peers
	installAction, err := chaincode.NewInstallAction(&chaincode.InstallCCArgs{chaincodeID, chaincodeVersion})
	if err != nil {
		appLogger.Debugf("NewInstallAction err [%v]\n", err)
		return err
	}
	if err = installAction.Execute(); err != nil {
		appLogger.Debugf("installAction err [%v]\n", err)
		return err
	}

	time.Sleep(time.Second * 2)

	//initialize chaincode on primary peer
	initAction, err := chaincode.NewInstantiateAction(&chaincode.InstantiateArgs{
		ChaincodeID: chaincodeID,
		Args:        args,
	})
	if err != nil {
		appLogger.Debugf("NewInstantiateAction err [%v]\n", err)
		return err
	}
	if _, err = initAction.Execute(); err != nil {
		appLogger.Debugf("initAction err [%v]\n", err)
		return err
	}
	return nil
}
