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

const (
	prefix = "/home/hxy/gopath/src/github.com/hyperledger/fabric-sdk-go/test"
)

func TestInstallAndInstantiateMarblesCC(t *testing.T) {
	testSetup := NewBaseSetupImpl(prefix)

	if err := testSetup.InstallAndInstantiateMarblesCC(); err != nil {
		t.Fatalf("InstallAndInstantiateMarblesCC return error: %v", err)
	}

	selftest, err := testSetup.read()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("selftest:", selftest)
}
