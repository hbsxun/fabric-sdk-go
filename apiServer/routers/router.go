package routers

import (
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/cert"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/chaincode"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/channel"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/query"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/user"
)

func init() {
	ns := beego.NewNamespace("/fabric",
		beego.NSNamespace("/enroll",
			beego.NSInclude(
				&cert.EnrollController{},
			),
		),
		beego.NSNamespace("/register",
			beego.NSInclude(
				&cert.RegisterController{},
			),
		),

		beego.NSNamespace("/install",
			beego.NSInclude(
				&chaincode.InstallCCController{},
			),
		),
		beego.NSNamespace("/instantiate",
			beego.NSInclude(
				&chaincode.InstantiateController{},
			),
		),
		beego.NSNamespace("/invokeCC",
			beego.NSInclude(
				&chaincode.InvokeController{},
			),
		),
		beego.NSNamespace("/queryCC",
			beego.NSInclude(
				&chaincode.QueryController{},
			),
		),
		beego.NSNamespace("/chaincodeInfo",
			beego.NSInclude(
				&chaincode.ChaincodeInfoController{},
			),
		),

		beego.NSNamespace("/createChannel",
			beego.NSInclude(
				&channel.ChannelCreateController{},
			),
		),
		beego.NSNamespace("/joinChannel",
			beego.NSInclude(
				&channel.ChannelJoinController{},
			),
		),

		beego.NSNamespace("/queryInstalled",
			beego.NSInclude(
				&query.QueryInstalledController{},
			),
		),
		beego.NSNamespace("/queryBlock",
			beego.NSInclude(
				&query.QueryBlockController{},
			),
		),
		beego.NSNamespace("/queryChannels",
			beego.NSInclude(
				&query.QueryChannelsController{},
			),
		),
		beego.NSNamespace("/queryBlockchainInfo",
			beego.NSInclude(
				&query.QueryInfoController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&user.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
