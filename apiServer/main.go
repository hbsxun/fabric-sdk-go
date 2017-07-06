package main

import (
	"github.com/astaxie/beego"
<<<<<<< HEAD
	_ "github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
=======
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/channel"
>>>>>>> 066d20b1ebc3d096375210849eb5667df1bb67c5
	_ "github.com/hyperledger/fabric-sdk-go/apiServer/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	channel.CreateAndJoinChannel()
	beego.Run()
}
