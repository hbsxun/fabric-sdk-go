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
	"errors"
	"fmt"

	fabricCaClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
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
	MaxEnrollments string      `json:"maxEnrollments"`
	Attributes     []Attribute `json:"attributes"`
}
type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type registerAction struct {
	common.ActionImpl
	req *fabricCaClient.RegistrationRequest
}

func NewRegisterAction(args *RegisterArgs) (*registerAction, error) {
	action, flags := &registerAction{}, &pflag.FlagSet{}

	var attrs []fabricCaClient.Attribute
	for _, attr := range args.Attributes {
		attrs = append(attrs, fabricCaClient.Attribute{
			Key:   attr.Key,
			Value: attr.Value,
		})
	}
	action.req = &fabricCaClient.RegistrationRequest{
		Name:           args.Name,
		Type:           "user",
		Affiliation:    "org1.department1",
		MaxEnrollments: 0,
		Attributes:     attrs,
	}

	err := action.Initialize(flags)
	return action, err
}

func (action *registerAction) Execute() (string, error) {

	fmt.Printf("Registering user [%s]\n", action.req.Name)
	caClient, err := fabricCaClient.NewFabricCAClient()
	if err != nil {
		return "", err
	}
	registrar, err := action.Client().LoadUserFromStateStore("admin")
	if err != nil {
		return "", err
	}
	secret, err := caClient.Register(registrar, action.req)
	if err != nil {
		return "", err
	} else {
		return secret, nil
	}

	return "", errors.New("Register User error")
}
