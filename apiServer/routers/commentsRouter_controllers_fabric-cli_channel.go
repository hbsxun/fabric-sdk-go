package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/channel:ChannelController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/channel:ChannelController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/CreateChannel`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/channel:ChannelController"] = append(beego.GlobalControllerRouter["github.com/hyperledger/fabric-sdk-go/apiServer/controllers/fabric-cli/channel:ChannelController"],
		beego.ControllerComments{
			Method: "JoinChannel",
			Router: `/JoinChannel`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
