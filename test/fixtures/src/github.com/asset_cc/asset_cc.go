package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"bytes"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("Fabric asset_cc chaincode")

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//定义描述资产的数据结构
type AssetInfo struct {
	ObjectType string `json:"doctype"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Price      int    `json:"price"`
	Owner      string `json:"owner"`
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function != "init" {
		return shim.Error("function name err, must be 'init'")
	}
	for i := 0; i < len(args); i++ {
		if len(args[i]) <= 0 {
			continue
		}
		var usrArgs [2]args

		usrArgs[0] = args[i]
		hash := sha256.New()
		hash.Write([]byte(usrArgs[0]))
		usrArgs[1] = base64.StdPadding.EncodeToString(hash.Sum(nil))

		t.AddUser(stub, usrArgs[:])
	}
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "AddUser" {
		return t.AddUser(stub, args)
	} else if function == "AddAsset" {
		return t.AddAsset(stub, args)
	} else if function == "Queryasset" {
		return t.Queryasset(stub, args)
	} else if function == "DeleteAsset" {
		return t.DeleteAsset(stub, args)
	} else if function == "UpdateAsset" {
		return t.UpdateAsset(stub, args)
	} else if function == "QueryAssetByOwner" {
		return t.QueryAssetByOwner(stub, args)
	} else if function == "GetAssetHistory" {
		return t.GetAssetHistory(stub, args)
	}
	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

//添加用户
func (t *SimpleChaincode) AddUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	fmt.Println("- start add user")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	name := args[0]
	encodedCertHash := args[1]

	userAsBytes, err := stub.GetState(name)
	if err != nil {
		return shim.Error("Failed to get user: " + err.Error())
	} else if userAsBytes != nil {
		fmt.Println("This user already exists: " + name)
		return shim.Error("This user already exists: " + name)
	}
	objectType := "USER"
	user := &UserInfo{objectType, name,
		encodedCertHash}
	userJSONasBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(name, userJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end add user")
	return shim.Success(nil)
}

//添加资产
func (t *SimpleChaincode) AddAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	//args[0:3] represent
	//Name Desc Price Owner
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	fmt.Println("- start add asset")
	if len(args[0]) <= 0 {
		return shim.Error("Name must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("Desc must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("Price must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("Owner argument must be a non-empty string")
	}

	price, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("price argument must be a numeric string")
	}

	//check if user exists

	userAsBytes, err := stub.GetState(userid)
	if err != nil {
		return shim.Error("Failed to get user: " + err.Error())
	} else if userAsBytes == nil {
		fmt.Println("This user does not exists: " + userid)
		return shim.Error("This user does not exists: " + userid)
	} else if string(userAsJson) != string(userAsBytes) {
		fmt.Println("This user does not match " + userid)
		return shim.Error("This user does not match " + userid)
	}
	assetAsBytes, err := stub.GetState(assetid)
	if err != nil {
		return shim.Error("Failed to get asset: " + err.Error())
	} else if assetAsBytes != nil {
		fmt.Println("This asset already exists: " + assetid)
		return shim.Error("This asset already exists: " + assetid)
	}
	objectType := "asset"
	asset := &AssetInfo{objectType, assetid, assetname, assetprice, user, owner}
	assetJSONasBytes, err := json.Marshal(asset)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(assetid, assetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end add asset")
	return shim.Success(nil)
}

//查询资产
func (t *SimpleChaincode) Queryasset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting id of the asset to query")
	}
	fmt.Println("- start query asset")
	id = args[0]
	valAsbytes, err := stub.GetState(id)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"asset does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}
	fmt.Println("- end query asset")
	return shim.Success(valAsbytes)
}

//删除资产
func (t *SimpleChaincode) DeleteAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var assetJSON AssetInfo
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	fmt.Println("- start delete asset")
	assetid := args[0]
	userid := args[1]
	username := args[2]
	user := UserInfo{"user", userid, username}
	valAsbytes, err := stub.GetState(assetid)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + assetid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"asset does not exist: " + assetid + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes),
		&assetJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + assetid + "\"}"
		return shim.Error(jsonResp)
	}
	if assetJSON.User != user {
		jsonResp = "{\"Error\":\"User does not match: " + userid + "\"}"
		return shim.Error(jsonResp)
	}
	err = stub.DelState(assetid)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	fmt.Println("- end delete asset")
	return shim.Success(nil)
}

func (t *SimpleChaincode) UpdateAsset(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var assetJSON AssetInfo
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	fmt.Println("- start update asset")
	assetid := args[0]
	updatename := args[1]
	updatecontent := args[2]
	userid := args[3]
	username := args[4]
	user := UserInfo{"user", userid, username}
	valAsbytes, err := stub.GetState(assetid)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + assetid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"asset does not exist: " + assetid + "\"}"
		return shim.Error(jsonResp)
	}
	err = json.Unmarshal([]byte(valAsbytes), &assetJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + assetid + "\"}"
		return shim.Error(jsonResp)
	}
	if assetJSON.User != user {
		jsonResp = "{\"Error\":\"User does not match: " + userid + "\"}"
		return shim.Error(jsonResp)
	}
	switch updatename {
	case "name":
		assetJSON.Name = updatecontent
	case "price":
		price, err := strconv.Atoi(updatecontent)
		if err != nil {
			return shim.Error(err.Error())
		}
		assetJSON.Price = price
	case "owner":
		assetJSON.Owner = updatecontent
	default:
		jsonResp = "{\"Error\":\"updatename is not exist: " + updatename + "\"}"
		return shim.Error(jsonResp)

	}

	assetJSONasBytes, err := json.Marshal(assetJSON)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(assetid, assetJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end update asset")
	return shim.Success(nil)
}

func (t *SimpleChaincode) QueryAssetByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	owner := strings.ToLower(args[0])

	queryString :=
		fmt.Sprintf("{\"selector\":{\"doctype\":\"asset\",\"owner\":\"%s\"}}", owner)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResultKey, queryResultRecord, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResultKey)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(
			string(queryResultRecord))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}
func (t *SimpleChaincode) GetAssetHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	assetid := args[0]

	fmt.Printf("- start GetAssetHistory: %s\n", assetid)

	resultsIterator, err := stub.GetHistoryForKey(assetid)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the model
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		txID, historicValue, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(txID)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// historicValue is a JSON model, so we write as-is
		buffer.WriteString(string(historicValue))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- GetAssetHistory returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
