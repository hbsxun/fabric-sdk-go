/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// ModelChaincode example simple Chaincode implementation
type ModelChaincode struct {
}

type model struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
	Desc       []byte `json:"desc"`
	Owner      string `json:"owner"`
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(ModelChaincode))
	if err != nil {
		fmt.Printf("Error starting Model chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *ModelChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *ModelChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("************************invoke is running " + function + "*********************************************")
	if len(args) < 1 {
		return shim.Error("The number of args must greater than 1. The first is  function name")
	}
	// Handle different functions
	if function == "initModel" { //create a new model
		return t.initModel(stub, args)
	} else if function == "transferModel" { //change owner of a specific model
		return t.transferModel(stub, args)
	} else if function == "delete" { //delete a model
		return t.delete(stub, args)
	} else if function == "readModel" { //read a model
		return t.readModel(stub, args)
	} else if function == "queryModelsByOwner" { //find models for owner X using rich query
		return t.queryModelsByOwner(stub, args)
	} else if function == "queryModels" { //find models based on an ad hoc rich query
		return t.queryModels(stub, args)
	} else if function == "getHistoryForModel" { //get history of values for a model
		return t.getHistoryForModel(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	//return shim.Error("Received unknown function invocation")
	s := fmt.Sprintf("function: %s\n, args: %v\n", function, args)
	return shim.Error(s)
}

// ============================================================
// initModel - create a new model, store into chaincode state
// ============================================================
func (t *ModelChaincode) initModel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	//   0       1       2
	// "Owner", "Name", "desc"
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init model")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	owner := strings.ToLower(args[0])
	modelName := args[1]
	desc := args[2]

	// ==== Check if model already exists ====
	modelAsBytes, err := stub.GetState(modelName)
	if err != nil {
		return shim.Error("Failed to get model: " + err.Error())
	} else if modelAsBytes != nil {
		fmt.Println("This model already exists: " + modelName)
		return shim.Error("This model already exists: " + modelName)
	}

	// ==== Create model object and marshal to JSON ====
	objectType := "model"
	model := &model{objectType, modelName, []byte(desc), owner}
	modelJSONasBytes, err := json.Marshal(model)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save model to state ===
	err = stub.PutState(modelName, modelJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// ===============================================
// readModel - read a model from chaincode state
// ===============================================
func (t *ModelChaincode) readModel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name, jsonResp string
	var err error

	fmt.Println("reading model: " + args[0])
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the model to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetState(name) //get the model from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Model does not exist: " + name + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ==================================================
// delete - remove a model key/value pair from state
// ==================================================
func (t *ModelChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	var modelJSON model
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	modelName := args[0]

	// to maintain the color~name index, we need to read the model first and get its color
	valAsbytes, err := stub.GetState(modelName) //get the model from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + modelName + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Model does not exist: " + modelName + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &modelJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + modelName + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(modelName) //remove the model from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	return shim.Success(nil)
}

// ===========================================================
// transfer a model by setting a new owner name on the model
// ===========================================================
func (t *ModelChaincode) transferModel(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0       1
	// "name", "newOwner"
	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	modelName := args[0]
	newOwner := strings.ToLower(args[1])
	fmt.Println("- start transferModel ", modelName, newOwner)

	modelAsBytes, err := stub.GetState(modelName)
	if err != nil {
		return shim.Error("Failed to get model:" + err.Error())
	} else if modelAsBytes == nil {
		return shim.Error("Model does not exist")
	}

	modelToTransfer := model{}
	err = json.Unmarshal(modelAsBytes, &modelToTransfer) //unmarshal it aka JSON.parse()
	if err != nil {
		return shim.Error(err.Error())
	}
	modelToTransfer.Owner = newOwner //change the owner

	modelJSONasBytes, _ := json.Marshal(modelToTransfer)
	err = stub.PutState(modelName, modelJSONasBytes) //rewrite the model
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end transferModel (success)")
	return shim.Success(nil)
}

// =======Rich queries =========================================================================
// Two examples of rich queries are provided below (parameterized query and ad hoc query).
// Rich queries pass a query string to the state database.
// Rich queries are only supported by state database implementations
//  that support rich query (e.g. CouchDB).
// The query string is in the syntax of the underlying state database.
// With rich queries there is no guarantee that the result set hasn't changed between
//  endorsement time and commit time, aka 'phantom reads'.
// Therefore, rich queries should not be used in update transactions, unless the
// application handles the possibility of result set changes between endorsement and commit time.
// Rich queries can be used for point-in-time queries against a peer.
// ============================================================================================

// ===== Example: Parameterized rich query =================================================
// queryModelsByOwner queries for models based on a passed in owner.
// This is an example of a parameterized query where the query logic is baked into the chaincode,
// and accepting a single query parameter (owner).
// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================================
func (t *ModelChaincode) queryModelsByOwner(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "bob"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	owner := strings.ToLower(args[0])

	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"model\",\"owner\":\"%s\"}}", owner)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// ===== Example: Ad hoc rich query ========================================================
// queryModels uses a query string to perform a query for models.
// Query string matching state database syntax is passed in and executed as is.
// Supports ad hoc queries that can be defined at runtime by the client.
// If this is not desired, follow the queryModelsForOwner example for parameterized queries.
// Only available on state databases that support rich query (e.g. CouchDB)
// =========================================================================================
func (t *ModelChaincode) queryModels(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0
	// "queryString"
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	queryString := args[0]

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
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
		buffer.WriteString(string(queryResultRecord))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (t *ModelChaincode) getHistoryForModel(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	modelName := args[0]

	fmt.Printf("- start getHistoryForModel: %s\n", modelName)

	resultsIterator, err := stub.GetHistoryForKey(modelName)
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

	fmt.Printf("- getHistoryForModel returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
