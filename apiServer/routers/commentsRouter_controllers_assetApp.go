package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "AddModel",
			Router: `/AddModel`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "DeleteModel",
			Router: `/DeleteModel/:ModelName`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "GetHistoryForModel",
			Router: `/GetHistoryForModel/:ModelName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "QueryModel",
			Router: `/QueryModel/:ModelName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "QueryModelsByOwner",
			Router: `/QueryModelsByOwner/:owner`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "TransferModel",
			Router: `/TransferModel`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:CertificateController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:CertificateController"],
		beego.ControllerComments{
			Method: "Enroll",
			Router: `/Enroll`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:CertificateController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:CertificateController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/Register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:InitializeController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:InitializeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "UpdatePasswd",
			Router: `/UpdatePasswd/:name/:oldPassword/:newPassword`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "VerifyUser",
			Router: `/VerifyUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/addUser`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "GetUserById",
			Router: `/getUserById/:userId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "GetUserByName",
			Router: `/getUserByName/:userName`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "UpdateInfo",
			Router: `/updateUser`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/userLogin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
