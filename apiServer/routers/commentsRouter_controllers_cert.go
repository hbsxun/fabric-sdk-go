package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/cert:CertificateController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/cert:CertificateController"],
		beego.ControllerComments{
			Method: "Enroll",
			Router: `/Enroll`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/cert:CertificateController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/cert:CertificateController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/Register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
