package main

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/hjwt"
	_ "github.com/hyperledger/fabric-sdk-go/apiServer/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//filter,  api Authorization
	beego.InsertFilter("/fabric/*", beego.BeforeRouter, func(ctx *context.Context) {
		//fmt.Println("URI:", ctx.Input.URI())
		//fmt.Println("Request:", ctx.Request.RequestURI)
		uri := ctx.Input.URI()
		if strings.HasPrefix(uri, "/fabric/user/addUser") || strings.HasPrefix(uri, "/fabric/user/userLogin") {
			return
		} else if strings.HasPrefix(uri, "/fabric/user/updateUser") || strings.HasPrefix(uri, "/fabric/user/getUserByName") || strings.HasPrefix(uri, "/fabric/user/getUserById") {
			token := ctx.Input.Cookie("Bearer")
			//token := ctx.Request.Header.Get("Authorization")
			fmt.Println(token)
			if valid, _ := hjwt.CheckToken(token); !valid {
				fmt.Println("Token is invalid or expiry")
				ctx.ResponseWriter.Write([]byte("Authorization failed, not login or you don't have the priviledge"))
				return
			}
		} else if strings.HasPrefix(uri, "/fabric/model/AddModel") || strings.HasPrefix(uri, "/fabric/model/DeleteModel") || strings.HasPrefix(uri, "/fabric/model/TransferModel") {
			token := ctx.Input.Cookie("Bearer")
			if _, isAdmin := hjwt.CheckToken(token); !isAdmin {
				fmt.Println("Don't have the 'admin' privilege")
				ctx.ResponseWriter.Write([]byte("Permission Denied, you don't have privilege."))
				return
			}
		}
	})
	beego.Run()
}
