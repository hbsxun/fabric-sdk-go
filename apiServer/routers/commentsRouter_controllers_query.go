package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/query:QueryInstalledController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/query:QueryInstalledController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:peerUrl`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
