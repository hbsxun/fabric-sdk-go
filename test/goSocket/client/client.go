package main

import (
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
	REGISTER   = "register"
	ENROLL     = "enroll"
	ADDASSET   = "addAsset"
	QUERYASSET = "queryAsset"
)

type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

func main() {
	server := "localhost:1024"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("connect success")
	//identity test
	IdentityTest(conn)
	//chaincode test

	//	ChaincodeTest(conn)
	//ledger test
	//LedgerTest(conn)
}
func ChaincodeTest(conn net.Conn) {
	txId, err := AddModel(conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("txId:", txId)

	modelId := "M1"
	modelInfo, err := QueryModel(modelId, conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("Owner#Name#Source#etc...:", modelInfo)
}

//QueryModel
func QueryModel(modelId string, conn net.Conn) (interface{}, error) {
	message := &Msg{
		Meta: map[string]interface{}{
			"meta": QUERYASSET,
			//"ID":        strconv.Itoa(i),
			"TimeStamp": time.Now().Format("2006-01-02 15:04:05"), //must be this time, 123456 2006
		},
		Content: map[string]interface{}{
			"name": "M1",
		},
	}
	//send
	result, _ := json.Marshal(message)
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
	if err != nil {
		return nil, fmt.Errorf("Receive data cannot be Unmarshal")
	}
	return recMap, nil

}

//AddModel
func AddModel(conn net.Conn) (interface{}, error) {
	message := &Msg{
		Meta: map[string]interface{}{
			"meta": ADDASSET,
			//"ID":        strconv.Itoa(i),
			"TimeStamp": time.Now().Format("2006-01-02 15:04:05"), //must be this time, 123456 2006
		},
		Content: map[string]interface{}{
			"owner":  "alice",
			"name":   "M1",
			"source": "Something....",
		},
	}
	//send
	result, _ := json.Marshal(message)
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
	if err != nil {
		return nil, fmt.Errorf("Receive data cannot be Unmarshal")
	}
	return recMap, nil

}

//IdentityTest user identity management
func IdentityTest(conn net.Conn) {
	enrollArgs, err := Register(conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Println("Name#Secret:", enrollArgs)

	certInfo, err := Enroll(enrollArgs, conn)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
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
			"meta":      ENROLL,
			"TimeStamp": time.Now().Format("2006-01-02 15:04:05"),
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
	//send
	result, _ := json.Marshal(message)
	conn.Write(utils.Enpack((result)))
	time.Sleep(1 * time.Second)

	//receive
	var buf = make([]byte, 4096)
	size, _ := conn.Read(buf)
	fmt.Println("receive: ", string(buf[:size]))

	//parse
	var recMap = make(map[string]interface{})
	err := json.Unmarshal(buf[:size], &recMap)
	if err != nil {
		return nil, fmt.Errorf("Receive data cannot be Unmarshal")
	}

	return recMap, nil
}

//Register user
func Register(conn net.Conn) (interface{}, error) {
	message := &Msg{
		Meta: map[string]interface{}{
			"meta": REGISTER,
			//"ID":        strconv.Itoa(i),
			"TimeStamp": time.Now().Format("2006-01-02 15:04:05"), //must be this time, 123456 2006
		},
		Content: map[string]interface{}{
			"name":           GetSession(),
			"type":           "user", //peer, app, user
			"maxEnrollments": 0,
			"affiliation":    "org1.department1",
			//"Attribute":["id","test"],
		},
	}
	//send
	result, _ := json.Marshal(message)
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
	if err != nil {
		return nil, fmt.Errorf("Receive data cannot be Unmarshal")
	}
	return recMap, nil
}

func GetSession() string {
	gs1 := time.Now().Unix()
	gs2 := strconv.FormatInt(gs1, 10)
	return gs2
}
