package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/hyperledger/fabric-ca/api"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
	"github.com/op/go-logging"
)

const ()

var logger = logging.MustGetLogger("Fabric-SDK-Server")

type Asset struct {
	Owner  string
	Name   string
	Source string
}
type RegisterResp struct {
	name, secret string
}
type EnrollResp struct {
	key, cert string
}

//RegisterUser
func RegisterUser(conn net.Conn) {
	remote := conn.RemoteAddr().String()
	for {
		buf := make([]byte, 512)
		size, err := conn.Read(buf)
		if err != nil {
			logger.Errorf("Read err [%s]", err)
			return
		}
		logger.Debugf("Receive from client [%s] is [%s]\n", remote, string(buf[:size]))
		//req RegistrationRequst: Name Type MaxEnrollments Affiliation Attribute
		var req fabricCAClient.RegistrationRequest
		err = json.Unmarshal(buf[:size], &req)
		if err != nil {
			logger.Errorf("Unmarshal err [%v]", err)
		}
		logger.Debugf("Unmarshal return:", req)

		//register user
		admin := sdkIgn.NewMember()
		nm, srt, err := admin.RegisterUser(req)
		if err != nil {
			logger.Errorf("registerUser err %v", err)
			return
		}

		//response secret
		resp := &RegisterResp{nm, srt}
		respJson, err := json.Marshal(resp)
		if err != nil {
			logger.Errorf("RegistrationResponse err %v", err)
		}
		conn.Write(respJson)
		conn.Close()
		break
	}
	logger.Debugf("RegisterUser success!")
}

//UserEnroll
func UserEnroll(conn net.Conn) {
	remote := conn.RemoteAddr().String()
	for {
		buf := make([]byte, 512)
		size, err := conn.Read(buf)
		if err != nil {
			logger.Errorf("Read err [%s]", err)
			return
		}
		logger.Debugf("Receive from client [%s] is [%s]\n", remote, string(buf[:size]))

		//req EnrollmentRequest: Name Secret Hosts Profile Label *CSR
		var req api.EnrollmentRequest
		err = json.Unmarshal(buf[:size], &req)
		if err != nil {
			logger.Errorf("Unmarshal err [%v]", err)
		}
		logger.Debugf("Unmarshal return:", req)

		//register user
		admin := sdkIgn.NewMember()
		key, cert, err := admin.UserEnroll(req.Name, req.Secret)
		//key, cert, err := admin.UserEnrollWithCSR(req)
		if err != nil {
			logger.Errorf("UserEnroll err %v", err)
			return
		}

		//response key cert
		resp := &EnrollResp{string(key), string(cert)}
		respJson, err := json.Marshal(resp)
		if err != nil {
			logger.Errorf("EnrollmentResponse err %v", err)
		}
		conn.Write(respJson)
		conn.Close()
		break
	}
	logger.Debugf("UserEnroll Success!")

}
func dealAsset(conn net.Conn) {
	remote := conn.RemoteAddr().String()
	fmt.Println(remote, "connected!")
	for {
		buf := make([]byte, 512)
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("Read err [%s]\n", err)
			return
		}
		fmt.Printf("Receive from client [%s] is [%s]\n", remote, string(buf[:size]))
		var model Model
		err = json.Unmarshal(buf[:size], &model)
		if err != nil {
			fmt.Printf("Unmarshal err [%v]", err)
		}
		fmt.Println("Model after Unmarshal:", model)

		conn.Write([]byte("Server has Receive"))
		conn.Close()
		break
	}
}

//handleClient
func handleClient(conn net.Conn) {
	//identity certificate

	//asset management

	//ledger query
}

func main() {
	fmt.Println("Starting the server")
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Printf("Listen failed [%v]", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept err [%s]\n", err)
			continue
		}
		go handleClient(conn)
	}
}
