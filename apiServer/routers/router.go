// @APIVersion 1.0.0
// @Title beego  API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact warm3snow@linux.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/cert"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/chaincode"
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers/query"
)

func init() {
	ns := beego.NewNamespace("/fabric",
		beego.NSNamespace("/enroll",
			beego.NSInclude(
				&cert.EnrollController{},
			),
		),
		beego.NSNamespace("/queryInstalled",
			beego.NSInclude(
				&query.QueryInstalledController{},
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
	)
	beego.AddNamespace(ns)
}
