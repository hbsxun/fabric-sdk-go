package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"bytes"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

//定义描述用户的数据结构
type userinfo struct {
	ObjectType string `json:"doctype"`
	Id         string `json:"id"`
	Name       string `json:"name"`
}

//定义描述资产的数据结构
type assertinfo struct {
	ObjectType string   `json:"doctype"`
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	User       userinfo `json:"user"`
	Owner      string   `json:"owner"`
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
	} else if function == "AddAssert" {
		return t.AddAssert(stub, args)
	} else if function == "QueryAssert" {
		return t.QueryAssert(stub, args)
	} else if function == "DeleteAssert" {
		return t.DeleteAssert(stub, args)
	} else if function == "UpdateAssert" {
		return t.UpdateAssert(stub, args)
	} else if function == "queryAssertByOwner" {
		return t.queryAssertByOwner(stub, args)
	} else if function == "getHistoryForAssert" {
		return t.getHistoryForAssert(stub, args)
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
	userid := strings.ToLower(args[0])
	username := strings.ToLower(args[1])
	userAsBytes, err := stub.GetState(userid)
	if err != nil {
		return shim.Error("Failed to get user: " + err.Error())
	} else if userAsBytes != nil {
		fmt.Println("This user already exists: " + userid)
		return shim.Error("This user already exists: " + userid)
	}
	objectType := "user"
	user := &userinfo{objectType, userid,
	username}
	userJSONasBytes, err := json.Marshal(user)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(userid, userJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end add user")
	return shim.Success(nil)
}

//添加资产
func (t *SimpleChaincode) AddAssert(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) < 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}
	fmt.Println("- start add assert")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	assertid := strings.ToLower(args[0])
	assertname := strings.ToLower(args[1])
	assertprice, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}
	userid := strings.ToLower(args[3])
	username := strings.ToLower(args[4])
	user := userinfo{"user", userid, username}
	owner := strings.ToLower(args[5])
	userAsJson, err := json.Marshal(&user)
	if err != nil {
		return shim.Error(err.Error())
	}
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
	assertAsBytes, err := stub.GetState(assertid)
	if err != nil {
		return shim.Error("Failed to get assert: " + err.Error())
	} else if assertAsBytes != nil {
		fmt.Println("This assert already exists: " + assertid)
		return shim.Error("This assert already exists: " + assertid)
	}
	objectType := "assert"
	assert := &assertinfo{objectType, assertid, assertname, assertprice, user, owner}
	assertJSONasBytes, err := json.Marshal(assert)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(assertid, assertJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end add assert")
	return shim.Success(nil)
}

//查询资产
func (t *SimpleChaincode) QueryAssert(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var id, jsonResp string
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting id of the assert to query")
	}
	fmt.Println("- start query assert")
	id = args[0]
	valAsbytes, err := stub.GetState(id)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"assert does not exist: " + id + "\"}"
		return shim.Error(jsonResp)
	}
	fmt.Println("- end query assert")
	return shim.Success(valAsbytes)
}

//删除资产
func (t *SimpleChaincode) DeleteAssert(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var assertJSON assertinfo
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	fmt.Println("- start delete assert")
	assertid := args[0]
	userid := args[1]
	username := args[2]
	user := userinfo{"user", userid, username}
	valAsbytes, err := stub.GetState(assertid)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + assertid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Assert does not exist: " + assertid + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes),
	&assertJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + assertid + "\"}"
		return shim.Error(jsonResp)
	}
	if assertJSON.User != user {
		jsonResp = "{\"Error\":\"User does not match: " + userid + "\"}"
		return shim.Error(jsonResp)
	}
	err = stub.DelState(assertid)
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}
	fmt.Println("- end delete assert")
	return shim.Success(nil)
}

func (t *SimpleChaincode) UpdateAssert(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var assertJSON assertinfo
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	fmt.Println("- start update assert")
	assertid := args[0]
	updatename := args[1]
	updatecontent := args[2]
	userid := args[3]
	username := args[4]
	user := userinfo{"user", userid, username}
	valAsbytes, err := stub.GetState(assertid)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + assertid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Assert does not exist: " + assertid + "\"}"
		return shim.Error(jsonResp)
	}
	err = json.Unmarshal([]byte(valAsbytes), &assertJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + assertid + "\"}"
		return shim.Error(jsonResp)
	}
	if assertJSON.User != user {
		jsonResp = "{\"Error\":\"User does not match: " + userid + "\"}"
		return shim.Error(jsonResp)
	}
	switch updatename {
	case "name":
		assertJSON.Name = updatecontent
	case "price":
		price, err := strconv.Atoi(updatecontent)
		if err != nil {
			return shim.Error(err.Error())
		}
		assertJSON.Price = price
	case "owner":
		assertJSON.Owner = updatecontent
	default:
		jsonResp = "{\"Error\":\"updatename is not exist: " + updatename + "\"}"
		return shim.Error(jsonResp)
		
	}

	assertJSONasBytes, err := json.Marshal(assertJSON)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(assertid, assertJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end update assert")
	return shim.Success(nil)
}

func (t *SimpleChaincode) queryAssertByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	owner := strings.ToLower(args[0])

	queryString :=
	fmt.Sprintf("{\"selector\":{\"doctype\":\"assert\",\"owner\":\"%s\"}}", owner)

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
func (t *SimpleChaincode) getHistoryForAssert(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	assertid := args[0]

	fmt.Printf("- start getHistoryForAssert: %s\n", assertid)

	resultsIterator, err := stub.GetHistoryForKey(assertid)
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

	fmt.Printf("- getHistoryForAssert returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
