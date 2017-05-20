package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/gislu/goSocket/client/utils"
)

const (
	REGISTER         = "register"
	ENROLL           = "enroll"
	ADDASSET         = "addAsset"
	QUERYASSET       = "queryAsset"
	QUERYBLOCK       = "queryBlock"       //on ledger
	QUERYTRANSACTION = "queryTransaction" //on ledger
)

type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

func main() {
	server := "localhost:8080"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	CheckError(err)

	fmt.Printf("connect [%s] success\n", server)
	//identity test
	//IdentityTest(conn)

	//chaincode test
	ChaincodeTest(conn)

	//ledger test
	//LedgerTest(conn)
}

func ChaincodeTest(conn net.Conn) {
	txId, err := AddModel(conn)
	CheckError(err)
	fmt.Println("txId:", txId)

	modelId := "M1"
	modelInfo, err := QueryModel(modelId, conn)
	CheckError(err)
	fmt.Println("Owner#Name#Source#etc...:", modelInfo)
}

//QueryModel
func QueryModel(modelId string, conn net.Conn) (interface{}, error) {
	message := &Msg{
		Meta: map[string]interface{}{
			"meta": QUERYASSET,
		},
		Content: map[string]interface{}{
			"name": modelId,
		},
	}
	return send(message, conn)
}

//AddModel
func AddModel(conn net.Conn) (interface{}, error) {
	message := &Msg{
		Meta: map[string]interface{}{
			"meta": ADDASSET,
		},
		Content: map[string]interface{}{
			"owner":  "alice",
			"name":   "asset1",
			"source": "Description",
		},
	}
	return send(message, conn)
}

//IdentityTest user identity management
func IdentityTest(conn net.Conn) {
	enrollArgs, err := Register(conn)
	CheckError(err)
	fmt.Println("Name#Secret:", enrollArgs)

	certInfo, err := Enroll(enrollArgs, conn)
	CheckError(err)
	fmt.Println("Key#Cert:", certInfo)

	defer conn.Close()
}

//Enroll user enroll(name, secret)
func Enroll(args interface{}, conn net.Conn) (interface{}, error) {
	argsMap, ok := args.(map[string]interface{})
	if !ok {
		return nil, errors.New("args should be a map")
	}
	if argsMap["name"] == nil || argsMap["secret"] == nil {
		return nil, errors.New("name or secret is nil")
	}
	message := &Msg{
		Meta: map[string]interface{}{
			"meta": ENROLL,
		},
		Content: map[string]interface{}{
			"name":    argsMap["name"],
			"secret":  argsMap["secret"],
			"hosts":   "www.example.com",
			"Profile": "profile...",
			"Label":   "for HSM",
			//"csr":     "*CSRInfo",
		},
	}
	return send(message, conn)
}

//Register user
func Register(conn net.Conn) (interface{}, error) {
	message := &Msg{
		Meta: map[string]interface{}{
			"meta": REGISTER,
		},
		Content: map[string]interface{}{
			"name":           GetSession(),
			"type":           "user", //peer, app, user
			"maxEnrollments": 0,
			"affiliation":    "org1.department1",
			//"Attribute":["id","test"],
		},
	}
	return send(message, conn)
}

//send
func send(message *Msg, conn net.Conn) (interface{}, error) {
	//send
	result, _ := json.Marshal(message)
	pack := utils.Enpack(result)
	/*
		fmt.Println("Send: ", string(pack))
		msgLen := utils.BytesToInt(pack[10:14])
		fmt.Println("Header+4+MsgLen: ", len(pack))
		fmt.Println("MessageLen: ", msgLen)
		fmt.Println("Message: ", string(pack[14:]))
	*/
	fmt.Println(hex.Dump(pack))

	conn.Write(utils.Enpack((result)))
	time.Sleep(1 * time.Second)
	//fmt.Println("send over")

	//receive
	var buf = make([]byte, 512)
	size, _ := conn.Read(buf)
	fmt.Println("receive: ", string(buf[:size]))

	//parse
	var recMap = make(map[string]interface{})
	err := json.Unmarshal(buf[:size], &recMap)
	CheckError(err)

	return recMap, err
}

//GetSession based-on time
func GetSession() string {
	gs1 := time.Now().Unix()
	gs2 := strconv.FormatInt(gs1, 10)
	return gs2
}

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
