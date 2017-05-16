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

import "testing"

func TestLedgerQueries(t *testing.T) {

	testSetup := NewBaseSetupImpl("/home/hxy/gopath/src/github.com/hyperledger/fabric-sdk-go/test")

	chain := testSetup.Chain
	//client := testSetup.Client
	accBook := NewLedger()

	//txId := "aff5c462b89d9ab2ec8eb86640a16fe1289aa8700f31fd252af05ae5375e1c3d"
	//accBook.QueryTrans(chain, txId)
	accBook.QueryBlock(chain)
	//accBook.QueryGenesisBlock(chain)
}
