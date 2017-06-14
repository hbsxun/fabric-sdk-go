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
	"github.com/hyperledger/fabric-sdk-go/apiServer/controllers"
)

func init() {
	ns := beego.NewNamespace("/fabric",
		beego.NSNamespace("/enroll",
			beego.NSInclude(
				&controllers.EnrollController{},
			),
		),
		beego.NSNamespace("/queryInstalled",
			beego.NSInclude(
				&controllers.QueryInstalledController{},
			),
		),
		beego.NSNamespace("/install",
			beego.NSInclude(
				&controllers.InstallCCController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
