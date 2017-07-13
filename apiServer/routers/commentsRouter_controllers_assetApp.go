package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "AddModel",
			Router: `/AddModel`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "QueryModel",
			Router: `/QueryModel/:ModelName`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "TransferModel",
			Router: `/TransferModel`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "QueryModelsByOwner",
			Router: `/QueryModelsByOwner/:owner`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "GetHistoryForModel",
			Router: `/GetHistoryForModel/:ModelName`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:AssetController"],
		beego.ControllerComments{
			Method: "DeleteModel",
			Router: `/DeleteModel/:ModelName`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:InitializeController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:InitializeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "Register",
			Router: `/addUser`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/userLogin`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/assetApp:UserManageController"],
		beego.ControllerComments{
			Method: "UpdateInfo",
			Router: `/updateUser`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

}
