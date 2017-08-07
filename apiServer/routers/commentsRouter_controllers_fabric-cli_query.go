package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query:QueryController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query:QueryController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/QueryBlock`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query:QueryController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query:QueryController"],
		beego.ControllerComments{
			Method: "QueryChainInfo",
			Router: `/QueryChainInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query:QueryController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query:QueryController"],
		beego.ControllerComments{
			Method: "QueryChannels",
			Router: `/QueryChannels`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query:QueryController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query:QueryController"],
		beego.ControllerComments{
			Method: "QueryInstalledChaincodes",
			Router: `/QueryInstalledChaincodes`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query:QueryController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/query:QueryController"],
		beego.ControllerComments{
			Method: "QueryTx",
			Router: `/QueryTx`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
