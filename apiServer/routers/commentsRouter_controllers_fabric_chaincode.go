package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric/chaincode:ChaincodeController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric/chaincode:ChaincodeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/ChaincodeInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric/chaincode:ChaincodeController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric/chaincode:ChaincodeController"],
		beego.ControllerComments{
			Method: "InstallCC",
			Router: `/InstallCC`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric/chaincode:ChaincodeController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric/chaincode:ChaincodeController"],
		beego.ControllerComments{
			Method: "InstantiateCC",
			Router: `/InstantiateCC`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric/chaincode:ChaincodeController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric/chaincode:ChaincodeController"],
		beego.ControllerComments{
			Method: "InvokeCC",
			Router: `/InvokeCC`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric/chaincode:ChaincodeController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric/chaincode:ChaincodeController"],
		beego.ControllerComments{
			Method: "QueryCC",
			Router: `/QueryCC`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
