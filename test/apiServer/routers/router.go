package routers

import (
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/chaincode",
			beego.NSNamespace("/asset",
				beego.NSInclude(
					&controllers.AssetController{},
				),
			),
		),
		beego.NSNamespace("/ca",
			beego.NSNamespace("/register",
				beego.NSInclude(
					&controllers.RegisterController{},
				),
			),
			beego.NSNamespace("/enroll",
				beego.NSInclude(
					&controllers.EnrollController{},
				),
			),
		),
	)
	beego.AddNamespace(ns)
}
