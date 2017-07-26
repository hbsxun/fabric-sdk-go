package routers

import (
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/chaincode"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/channel"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query"
	//"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/channel"
)

func init() {
	ns := beego.NewNamespace("/fabric",
		/*beego.NSNamespace("/enroll",
			beego.NSInclude(
				&cert.EnrollController{},
			),
		),
		beego.NSNamespace("/register",
			beego.NSInclude(
				&cert.RegisterController{},
			),
		),*/
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
		),*/
	)
	beego.AddNamespace(ns)
}
