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

func TestLedgerQueries(t *testing.T) {

	testSetup := NewBaseSetupImpl("/home/hxy/gopath/src/github.com/hyperledger/fabric-sdk-go/test")

	chain := testSetup.Chain
	//client := testSetup.Client
	accBook := NewLedger(chain)

	/*
		txId := "9f3de24ca5ba728db5e902d09dd472c589921b2aae996358a5886e5c6da6a137"
		txInfo, err := accBook.QueryTrans(txId)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(txInfo)
	*/
	blocks, err := accBook.QueryBlocks()
	if err != nil {
		t.Fatalf("QueryBlocks failed [%s]\n", err)
	}
	fmt.Println(blocks)

	genesisBlock, err := accBook.QueryBlockByNumber(0)
	if err != nil {
		t.Fatalf("Query Genesis Block failed [%s]\n", err)
	}
	fmt.Println(genesisBlock)
}
