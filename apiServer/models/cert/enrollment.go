/*
Copyright Beijing Sansec Technology Development Co., Ltd. All Rights Reserved.

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
package cert

import (
	"encoding/base64"
	"fmt"

	"github.com/hyperledger/fabric-ca/api"
	fabricCaClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	"github.com/spf13/pflag"
)

type EnrollReq api.EnrollmentRequest

type enrollAction struct {
	common.ActionImpl
}

func NewEnrollAction(flags *pflag.FlagSet) (*enrollAction, error) {
	action := &enrollAction{}
	err := action.Initialize(flags)
	return action, err
}

func (action *enrollAction) Enroll(req EnrollReq) (string, error) {
	fmt.Printf("enrolling user [%s]\n", req.Name)
	caClient, err := fabricCaClient.NewFabricCAClient()
	if err != nil {
		return "", err
	}
	key, cert, err := caClient.Enroll(req.Name, req.Secret)
	if err != nil {
		return "", err
	} else {
		return fmt.Sprintf("{Key: %s \nCert: %s}\n", base64.StdEncoding.EncodeToString(key), base64.StdEncoding.EncodeToString(cert)), nil
	}

	return "", nil
}
