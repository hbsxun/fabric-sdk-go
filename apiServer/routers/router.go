package routers

import (
	"github.com/astaxie/beego"
<<<<<<< HEAD
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/chaincode"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/channel"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query"
	//"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/channel"
=======
	//"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/cert"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/chaincode"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/channel"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query"
>>>>>>> upstream/v1.0.0
)

func init() {
	ns := beego.NewNamespace("/fabric",
<<<<<<< HEAD
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
=======
		beego.NSNamespace("/cert",
			beego.NSInclude(
				&cert.CertController{},
			),
		),

>>>>>>> upstream/v1.0.0
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
<<<<<<< HEAD
=======

>>>>>>> upstream/v1.0.0
		beego.NSNamespace("/query",
			beego.NSInclude(
				&query.QueryController{},
			),
		),
<<<<<<< HEAD

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
=======
		/*
			beego.NSNamespace("/user",
				beego.NSInclude(
					&user.UserController{},
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
		*/
>>>>>>> upstream/v1.0.0
	)
	beego.AddNamespace(ns)
}
