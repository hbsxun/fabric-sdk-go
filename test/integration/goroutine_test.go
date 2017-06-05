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
	"log"
	"os"
	"testing"
)

var testSetup BaseSetupImpl

func TestMain(m *testing.M) {
	testSetup = BaseSetupImpl{
		ConfigFile:      "../fixtures/config/config_test.yaml",
		ChainID:         "testchannel",
		ChannelConfig:   "../fixtures/channel/testchannel.tx",
		ConnectEventHub: true,
	}

	if err := testSetup.Initialize(); err != nil {
		log.Fatalf(err.Error())
	}

	if err := testSetup.InstallAndInstantiateExampleCC(); err != nil {
		log.Fatalf("InstallAndInstantiateExampleCC return error: %v", err)
	}

	//runtime.GOMAXPROCS(runtime.NumCPU())

	log.Println("[TestMain] before run()")
	code := m.Run()
	log.Println("[TestMain] after run()")
	os.Exit(code)
}

func BenchmarkQuery(b *testing.B) {
	b.SetParallelism(200)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testSetup.QueryAsset()
		}
	})
}

func BenchmarkInvoke(b *testing.B) {
	b.SetParallelism(200)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testSetup.MoveFunds()
		}
	})
}
