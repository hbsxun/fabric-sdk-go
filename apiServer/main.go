package main

import (
	"github.com/astaxie/beego"
	//"github.com/hyperledger/fabric-sdk-go/apiServer/models/channel"
	//_ "github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
	_ "github.com/hyperledger/fabric-sdk-go/apiServer/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//channel.CreateAndJoinChannel()
	beego.Run()
}
