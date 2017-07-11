package AssetApp

import (
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/chaincode"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
	"github.com/op/go-logging"
)

var appLogger = logging.MustGetLogger("AssetApp")

type UserManager interface {
	Register(*user.User) bool
	Login(name, passwd string) bool
	UpdateInfo(*user.User) bool
}

type Chaincode interface {
	//admin
	Deploy(chaincodeName, chaincodeVersion string) error    //install chaincode to all peers
	Instantiate(*chaincode.InstantiateArgs) (string, error) //returns a txid
	RegisterAsset(*chaincode.InvokeArgs) (string, error)    //returns a txid
	DeleteAsset(*chaincode.InvokeArgs) (string, error)      //returns a txid
	RegisterUser(*chaincode.InvokeArgs) (string, error)     //returns a txid

	//common
	QueryAssetByName(assetName string) ([]string, error)

	//user
	TransferAsset(assetName, newOwner string) (string, error) //returns a txid
	QueryAssetByOwner(owner string) ([][]string, error)
}

type Certificate interface {
	GetIdentity(string) (string, string) //name --> key, cert
	SaveToDB(string, string, string) error
}
