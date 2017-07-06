package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"],
		beego.ControllerComments{
			Method: "Adduser",
			Router: `/Adduser`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"],
		beego.ControllerComments{
			Method: "Addassert",
			Router: `/Addassert`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"],
		beego.ControllerComments{
			Method: "Queryassert",
			Router: `/Queryassert/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"],
		beego.ControllerComments{
			Method: "Querybyowner",
			Router: `/Querybyowner/:owner`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"],
		beego.ControllerComments{
			Method: "Gethistoryforassert",
			Router: `/gethistoryforassert/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"],
		beego.ControllerComments{
			Method: "Updateassert",
			Router: `/Updateassert`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/AssetApp:AssetController"],
		beego.ControllerComments{
			Method: "Deleteassert",
			Router: `/Deleteassert`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

}
