/*
Copyright SecureKey Technologies Inc. All Rights Reserved.


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at


      http://www.apache.org/licenses/LICENSE-2.0


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package integration

import (
	"fmt"
	"testing"
)

func TestModelCCInvoke(t *testing.T) {

	/*
		testSetup := BaseSetupImpl{
			ConfigFile:      "../fixtures/config/config_test.yaml",
			ChainID:         "testchannel",
			ChannelConfig:   "../fixtures/channel/testchannel.tx",
			ConnectEventHub: true,
		}

		if err := testSetup.Initialize(); err != nil {
			t.Fatalf(err.Error())
		}
	*/
	testSetup := NewBaseSetupImpl("/home/hxy/gopath/src/github.com/hyperledger/fabric-sdk-go/test")

	if err := testSetup.InstallAndInstantiateModelCC(); err != nil {
		t.Fatalf("InstallAndInstantiateModelCC return error: %v", err)
	}

	//add Model
	model := &Model{
		Owner:  "Alice",
		Name:   "M1",
		Source: "blabla",
	}
	_, err := testSetup.AddModel(model)
	if err != nil {
		t.Fatalf("Add Model return error: %v", err)
		return
	}

	//query Model
	modelInfo, err := testSetup.QueryModel(model.Name)
	if err != nil {
		t.Errorf("getModel info return error: %v", err)
		return
	}
	fmt.Printf("***Model info: %s\n", modelInfo)

	//transfer Model
	_, err = testSetup.TransferModel()
	if err != nil {
		t.Fatalf("TransferModel return error: %v", err)
		return
	}

	//query Model
	modelInfo, err = testSetup.QueryModelByOwner("alice")
	if err != nil {
		t.Errorf("getModel info return error: %v", err)
		return
	}
	fmt.Printf("***Model info after transfer: %s\n", modelInfo)
	modelInfo, err = testSetup.QueryModelByOwner("bob")
	if err != nil {
		t.Errorf("getModel info return error: %v", err)
		return
	}
	fmt.Printf("***Model info after transfer: %s\n", modelInfo)

	//query Model history
	history, err := testSetup.GetHistoryForModel()
	if err != nil {
		t.Fatalf("GetHistoryForModel return error: %v", err)
		return
	}
	fmt.Printf("***Model history:\n %v \n*********\n", history)
}
