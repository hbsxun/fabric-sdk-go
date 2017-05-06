package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInit(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}
func checkState(t *testing.T, stub *shim.MockStub, name string, value string) {
	bytes := stub.State[name]
	if bytes == nil {
		fmt.Println("State", name, "failed to get value")
		t.FailNow()
	}
	if string(bytes) != value {
		fmt.Println("State:", name, "!=", value)
		t.FailNow()
	}
}
func checkQuery(t *testing.T, stub *shim.MockStub, name string, value string) {
	res := stub.MockInvoke("1", [][]byte{[]byte("readModel"), []byte(name)})
	if res.Status != shim.OK {
		fmt.Println("Query failed:", name, string(res.String()))
		t.FailNow()
	}
	if res.Payload == nil {
		fmt.Println("Query failed:", name)
		t.FailNow()
	}
	if string(res.Payload) != value {
		fmt.Println("Query:", name, "is", string(res.Payload), " != ", value)
		t.FailNow()
	}
	fmt.Println("State:", string(res.Payload))
}
func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		fmt.Println("Invoke", args, "failed", string(res.Message))
		t.FailNow()
	}
}
func Test_Init(t *testing.T) {
	scc := new(ModelChaincode)
	stub := shim.NewMockStub("ModelCC", scc)
	checkInit(t, stub, [][]byte{})
}
func Test_Invoke(t *testing.T) {
	scc := new(ModelChaincode)
	stub := shim.NewMockStub("ModelCC", scc)

	//	checkInit(t, stub, nil)
	checkInvoke(t, stub, [][]byte{
		[]byte("initModel"),
		[]byte("programer1"),
		[]byte("match-1"),
		[]byte("Java"),
	})
	checkInvoke(t, stub, [][]byte{
		[]byte("initModel"),
		[]byte("programer2"),
		[]byte("match-2"),
		[]byte("Go"),
	})
	checkInvoke(t, stub, [][]byte{
		[]byte("initModel"),
		[]byte("programer1"),
		[]byte("match-1-2"),
		[]byte("C++"),
	})
	model := &model{
		Name:       "match-1",
		ObjectType: "model",
		Owner:      "programer1",
		Source:     []byte("Java"),
	}
	modelAsJson, err := json.Marshal(model)
	if err != nil {
		t.Error(err.Error())
		t.FailNow()
	}
	checkQuery(t, stub, "match-1", string(modelAsJson))
	//Mock not implemented
	/*
		checkInvoke(t, stub, [][]byte{
			[]byte("queryModelsByOwner"),
			[]byte("programer1"),
		})
	*/
}
