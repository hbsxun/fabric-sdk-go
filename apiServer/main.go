package main

import (
	"os"

	"github.com/astaxie/beego"
	//"github.com/hyperledger/fabric-sdk-go/apiServer/models/channel"
	//_ "github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
	_ "github.com/hyperledger/fabric-sdk-go/apiServer/routers"
	"github.com/hyperledger/fabric-sdk-go/fabric-cli/maincmd"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	if maincmd.NewFabricCLICmd().Execute() != nil {
		os.Exit(1)
	}

	//channel.CreateAndJoinChannel()
	beego.Run()
}
