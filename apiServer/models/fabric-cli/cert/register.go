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
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	ca "github.com/hyperledger/fabric-sdk-go/api/apifabca"
	"github.com/hyperledger/fabric-sdk-go/api/apifabclient"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/config"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/pkg/fabric-ca-client"
	fabricClient "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/identity"
	kvs "github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/keyvaluestore"
	bccspFactory "github.com/hyperledger/fabric/bccsp/factory"
	"github.com/spf13/pflag"
)

/*
const (
	RegisterNameFlag           = "Name"
	RegisterTypeFlag           = "Type"
	RegisterMaxEnrollmentsFlag = "MaxEnrollments"
	RegisterAffiliationFlag    = "Affiliation"
	RegisterAttributesFlag     = "Attributes"
)

	flags := certRegisterCmd.Flags()
	flags.StringVar(&registerReq.Name, RegisterNameFlag, registerReq.Name, "The user name for register")
	flags.StringVar(&registerReq.Type, RegisterTypeFlag, registerReq.Type, "The type of the user to register, should be user/app/peer")
	flags.StringVar(&registerReq.Affiliation, RegisterAffiliationFlag, registerReq.Affiliation, "The Org where the user affiliated")
	flags.IntVar(&registerReq.MaxEnrollments, RegisterMaxEnrollmentsFlag, registerReq.MaxEnrollments, "The max time for user to enroll, '0' means infinite times")
	//flags.StringArrayVar(&registerReq.Attributes, RegisterAttributesFlag, registerReq.Attributes, "The attributes which user want to use for fulture, it's a key-value arrays, Example {'k1':'v1','k2':'v2'}")
*/
type RegisterArgs struct {
	Name           string      `json:"name"`
	Type           string      `json:"type"`
	Affiliation    string      `json:"affiliation"`
	CAName         string      `json:"caName"`
	MaxEnrollments string      `json:"maxEnrollments"`
	Attributes     []Attribute `json:"attributes"`
}
type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type registerAction struct {
	common.Action
	req *ca.RegistrationRequest
}

func NewRegisterAction(args *RegisterArgs) (*registerAction, error) {
	action, flags := &registerAction{}, &pflag.FlagSet{}

	/*
		var attrs []fabricCAClient.Attribute
		for _, attr := range args.Attributes {
			attrs = append(attrs, fabricCAClient.Attribute{
				Key:   attr.Key,
				Value: attr.Value,
			})
		}
	*/
	if args.Name == "" {
		return nil, errors.New("Must specify a user name")
	}
	if args.Type != "user" && args.Type != "admin" {
		args.Type = "user"
	}
	if args.Affiliation == "" {
		args.Affiliation = "org1.department1"
	}
	if args.CAName == "" {
		args.CAName = "ca-org1"
	}
	action.req = &ca.RegistrationRequest{
		Name:           args.Name,
		Type:           args.Type,
		Affiliation:    args.Affiliation,
		CAName:         args.CAName,
		MaxEnrollments: 0,
		//Attributes: attrs,
	}

	err := action.Initialize(flags)
	return action, err
}

func (action *registerAction) Execute() (string, error) {
	fmt.Printf("Registering user [%s]\n", action.req.Name)

	var fabricCAConfig apiconfig.Config
	fabricCAConfig, err := config.InitConfig(common.Config().ConfigFile())
	if err != nil {
		return "", err
	}
	//client := fabricClient.NewClient(fabricCAConfig)
	caClient, err := fabricCAClient.NewFabricCAClient(fabricCAConfig, common.Config().OrgID())
	if err != nil {
		return "", err
	}

	//Note 1 ********************************
	//Default return peerorg1User1, role for register should be admin of orgx
	/*
		adminUser, err := action.Client().LoadUserFromStateStore("admin")
		fmt.Printf("adminUser: %v\n", adminUser)
		fmt.Printf("req: %v\n", action.req)
		if err != nil {
			return "", err
		}
	*/
	//Note 2**********************************
	//fabric-cli use pre-enrolled "Admin@org1.example.com",  fabric-sdk-go enroll a "admin".
	/*
		adminUser := action.OrgAdminUser(common.Config().OrgID())
		secret, err := caClient.Register(adminUser, action.req)
	*/
	adminUser, err := loadAdminUser(fabricCAConfig, common.Config().OrgID())
	if err != nil {
		return "", err
	}
	fmt.Printf("adminUser: %v\n", adminUser)
	fmt.Printf("req: %v\n", action.req)
	secret, err := caClient.Register(adminUser, action.req)
	if err != nil {
		return "", err
	} else {
		println("secret:", secret)
		return secret, nil
	}

	return "", errors.New("Register User error")
}

func loadAdminUser(testFabricCAConfig apiconfig.Config, orgName string) (adminUser apifabclient.User, err error) {
	mspID, err := testFabricCAConfig.MspID(orgName)
	if err != nil {
		return nil, fmt.Errorf("GetMapId() return error: %v", err)
	}
	client := fabricClient.NewClient(testFabricCAConfig)

	err = bccspFactory.InitFactories(testFabricCAConfig.CSPConfig())
	if err != nil {
		return nil, fmt.Errorf("Failed getting ephemeral software-based BCCSP [%s]", err)
	}

	cryptoSuite := bccspFactory.GetDefault()

	client.SetCryptoSuite(cryptoSuite)
	stateStore, err := kvs.CreateNewFileKeyValueStore("/tmp/enroll_user")
	if err != nil {
		return nil, fmt.Errorf("CreateNewFileKeyValueStore return error[%s]", err)
	}
	client.SetStateStore(stateStore)

	caClient, err := fabricCAClient.NewFabricCAClient(testFabricCAConfig, orgName)
	if err != nil {
		return nil, fmt.Errorf("NewFabricCAClient return error: %v", err)
	}

	// Admin user is used to register, enrol and revoke a test user
	adminUser, err = client.LoadUserFromStateStore("admin")

	if err != nil {
		return nil, fmt.Errorf("client.LoadUserFromStateStore return error: %v", err)
	}
	if adminUser == nil {
		key, cert, err := caClient.Enroll("admin", "adminpw")
		if err != nil || key == nil || cert == nil {
			return nil, fmt.Errorf("Enroll failed or return error")
		}
		certPem, _ := pem.Decode(cert)
		if err != nil {
			return nil, fmt.Errorf("pem Decode return error: %v", err)
		}

		cert509, err := x509.ParseCertificate(certPem.Bytes)
		if err != nil {
			return nil, fmt.Errorf("x509 ParseCertificate return error: %v", err)
		}
		if cert509.Subject.CommonName != "admin" {
			return nil, fmt.Errorf("CommonName in x509 cert is not the enrollmentID")
		}
		adminUser2 := identity.NewUser("admin", mspID)
		adminUser2.SetPrivateKey(key)
		adminUser2.SetEnrollmentCertificate(cert)
		err = client.SaveUserToStateStore(adminUser2, false)
		if err != nil {
			return nil, fmt.Errorf("client.SaveUserToStateStore return error: %v", err)
		}
		adminUser, err = client.LoadUserFromStateStore("admin")
		if err != nil {
			return nil, fmt.Errorf("client.LoadUserFromStateStore return error: %v", err)
		}
		if adminUser == nil {
			return nil, fmt.Errorf("client.LoadUserFromStateStore return nil")
		}
	}
	return adminUser, nil
}
