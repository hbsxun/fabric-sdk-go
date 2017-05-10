package main

import (
	"encoding/json"
	"fmt"
	"net"

	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("Fabric-SDK-Server")

type Model struct {
	Owner  string
	Name   string
	Source string
}
type Enrollment struct {
	name, secret string
}

func registerUser(conn net.Conn) {
	remote := conn.RemoteAddr().String()
	for {
		buf := make([]byte, 512)
		size, err := conn.Read(buf)
		if err != nil {
			logger.Debugf("Read err [%s]", err)
			return
		}
		logger.Debugf("Receive from client [%s] is [%s]\n", remote, string(buf[:size]))

		var req fabricCAClient.RegistrationRequest
		err = json.Unmarshal(buf[:size], &req)
		if err != nil {
			logger.Debugf("Unmarshal err [%v]", err)
		}
		logger.Debugf("Unmarshal return:", req)

		admin := sdkIgn.NewMember()
		id, otp, err := admin.RegisterUser(req)
		if err != nil {
			logger.Debugf("registerUser err %v", err)
			return
		}

		resp := &Enrollment{name, secret}

		conn.Write([]byte("Response OK!"))
		conn.Close()
		break
	}

}
func dealModel(conn net.Conn) {
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
		go dealModel(conn)
	}
}
