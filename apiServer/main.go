package main

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/assetApp/auth"
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
		uri := ctx.Input.URI()
		if strings.HasPrefix(uri, "/fabric/user/addUser") || strings.HasPrefix(uri, "/fabric/user/userLogin") {
			return
		} else if strings.Contains(uri, "/user/") || strings.Contains(uri, "/cert/") { //*************should login
			token := ctx.Input.Cookie("Bearer")
			fmt.Println(token)

			if valid := auth.IsTokenValid(token); !valid {
				fmt.Println("Token is invalid or expiry")
				ctx.ResponseWriter.Write([]byte("Login first please :-)"))
				return
			}
		} else if strings.Contains(uri, "/model/") { //********************should enroll
			_, name := auth.GetIdAndName(ctx.Input.Cookie("Bearer"))
			enrollments := ctx.Input.Cookie(name)
			identity, err := auth.UnSerialize(enrollments)
			_ = identity
			if err != nil {
				fmt.Println("Don't have a valid Ecert, please go to `enroll`")
				ctx.ResponseWriter.Write([]byte("Enroll first to get your certificate :-()"))
				return
			}
			if strings.HasPrefix(uri, "/fabric/model/AddModel") || strings.HasPrefix(uri, "/fabric/model/DeleteModel") {
				token := ctx.Input.Cookie("Bearer")
				if isAdmin := auth.IsAdmin(token); !isAdmin {
					fmt.Println("AddModel and DeleteModel operations should be executed by 'admin'")
					ctx.ResponseWriter.Write([]byte("Sorry, you don't have the priviledge to do that. 0(^o^)0"))
					return
				}
			} else if strings.HasPrefix(uri, "/fabric/model/TransferModel") {
				token := ctx.Input.Cookie("Bearer")
				if isAdmin := auth.IsAdmin(token); isAdmin {
					fmt.Println("Only user himself/herself can TransferModel their own Asset")
					ctx.ResponseWriter.Write([]byte("Only user himself/herself can TransferModel their own Asset, (*^_^*)"))
					return
				}
			}
		}
	})
	beego.Run()
}
