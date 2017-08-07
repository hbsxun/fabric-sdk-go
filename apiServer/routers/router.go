package routers

import (
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/cert"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/chaincode"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/channel"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query"
)

func init() {
	ns := beego.NewNamespace("/fabric",
		beego.NSNamespace("/cert",
			beego.NSInclude(
				&cert.CertController{},
			),
		),
		beego.NSNamespace("/channel",
			beego.NSInclude(
				&channel.ChannelController{},
			),
		),
		beego.NSNamespace("/chaincode",
			beego.NSInclude(
				&chaincode.ChaincodeController{},
			),
		),
		beego.NSNamespace("/query",
			beego.NSInclude(
				&query.QueryController{},
			),
		),
		/*beego.NSNamespace("/user",
			beego.NSInclude(
				&user.UserController{},
			),
		),*/

		beego.NSNamespace("/user",
			beego.NSInclude(
				&assetApp.UserManageController{},
			),
		),
		beego.NSNamespace("/initialize",
			beego.NSInclude(
				&assetApp.InitializeController{},
			),
		),
		beego.NSNamespace("/model",
			beego.NSInclude(
				&assetApp.AssetController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
