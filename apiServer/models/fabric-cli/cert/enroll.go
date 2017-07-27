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
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/config"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/pkg/fabric-ca-client"
	fabricClient "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/identity"
	kvs "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/keyvaluestore"
	bccspFactory "github.com/hyperledger/fabric/bccsp/factory"
	"github.com/spf13/pflag"
)

type EnrollArgs struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

type enrollAction struct {
	common.Action
	req *EnrollArgs
}

func NewEnrollAction(req *EnrollArgs) (*enrollAction, error) {
	action := &enrollAction{
		req: req,
	}

	flags := &pflag.FlagSet{}
	err := action.Initialize(flags)
	return action, err
}

func (action *enrollAction) Execute() (string, error) {
	fmt.Printf("enrolling user [%s]\n", action.req.Name)

	//ca config
	var testFabricCAConfig apiconfig.Config
	testFabricCAConfig, err := config.InitConfig(common.Config().ConfigFile())
	if err != nil {
		return "", err
	}

	orgName := common.Config().OrgID()
	mspID, err := testFabricCAConfig.MspID(orgName)
	if err != nil {
		return "", fmt.Errorf("GetMapId() return error: %v", err)
	}
	client := fabricClient.NewClient(testFabricCAConfig)

	err = bccspFactory.InitFactories(testFabricCAConfig.CSPConfig())
	if err != nil {
		return "", fmt.Errorf("Failed getting ephemeral software-based BCCSP [%s]", err)
	}

	cryptoSuite := bccspFactory.GetDefault()

	client.SetCryptoSuite(cryptoSuite)
	stateStore, err := kvs.CreateNewFileKeyValueStore("/tmp/enroll_user")
	if err != nil {
		return "", fmt.Errorf("CreateNewFileKeyValueStore return error[%s]", err)
	}
	client.SetStateStore(stateStore)

	caClient, err := fabricCAClient.NewFabricCAClient(testFabricCAConfig, orgName)
	if err != nil {
		return "", fmt.Errorf("NewFabricCAClient return error: %v", err)
	}

	key, cert, err := caClient.Enroll(action.req.Name, action.req.Secret)
	if err != nil {
		return "", err
	}
	//check if cert is this user's
	certPem, _ := pem.Decode(cert)
	if err != nil {
		return "", fmt.Errorf("pem Decode return error: %v", err)
	}

	cert509, err := x509.ParseCertificate(certPem.Bytes)
	if err != nil {
		return "", fmt.Errorf("x509 ParseCertificate return error: %v", err)
	}
	if cert509.Subject.CommonName != action.req.Name {
		return "", fmt.Errorf("CommonName in x509 cert is not the enrollmentID")
	}

	enrolledUser := identity.NewUser(action.req.Name, mspID)
	enrolledUser.SetPrivateKey(key)
	enrolledUser.SetEnrollmentCertificate(cert)
	err = client.SaveUserToStateStore(enrolledUser, false)
	if err != nil {
		return "", fmt.Errorf("client.SaveUserToStateStore return error: %v", err)
	}

	return base64.StdEncoding.EncodeToString(key.SKI()) + "." + base64.StdEncoding.EncodeToString(cert), nil
}
