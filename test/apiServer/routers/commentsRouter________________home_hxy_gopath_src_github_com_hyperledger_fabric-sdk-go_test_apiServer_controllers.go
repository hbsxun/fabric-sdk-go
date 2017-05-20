package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:AssetController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:EnrollController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/test/apiServer/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
