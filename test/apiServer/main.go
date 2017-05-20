package main

import (
	_ "github.com/hyperledger/fabric-sdk-go/test/apiServer/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

