package routers

import (
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp"
	//	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabricca"
)

func init() {
	ns := beego.NewNamespace("/fabric",
		/*
			beego.NSNamespace("/cert",
				beego.NSInclude(
					&fabricca.CertController{},
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
		*/

		beego.NSNamespace("/cert",
			beego.NSInclude(
				&assetApp.CertificateController{},
			),
		),
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
