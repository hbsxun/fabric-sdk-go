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
	"fmt"

	fabricCaClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	RegisterNameFlag           = "Name"
	RegisterTypeFlag           = "Type"
	RegisterMaxEnrollmentsFlag = "MaxEnrollments"
	RegisterAffiliationFlag    = "Affiliation"
	RegisterAttributesFlag     = "Attributes"
)

var registerReq *fabricCaClient.RegistrationRequest

var certRegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Register User",
	Long:  "Register a new user before enroll for key and cert",
	Run: func(cmd *cobra.Command, args []string) {
		if registerReq.Name == "" {
			fmt.Printf("\nMust specify the Registration Name\n\n")
			cmd.HelpFunc()(cmd, args)
			return
		}
		action, err := newRegisterAction(cmd.Flags())
		if err != nil {
			common.Logger.Criticalf("Error while initializing installAction: %v", err)
			return
		}
		err = action.register()
		if err != nil {
			common.Logger.Criticalf("Error while running installAction: %v", err)
			return
		}
	},
}

func getCertRegisterCmd() *cobra.Command {
	registerReq = &fabricCaClient.RegistrationRequest{
		Type:           "user",
		Affiliation:    "org1.department1",
		MaxEnrollments: 0,
		Attributes: []fabricCaClient.Attribute{
			fabricCaClient.Attribute{
				Key:   "key1",
				Value: "Trade",
			},
			fabricCaClient.Attribute{
				Key:   "key2",
				Value: "Store",
			},
		},
	}
	flags := certRegisterCmd.Flags()
	flags.StringVar(&registerReq.Name, RegisterNameFlag, registerReq.Name, "The user name for register")
	flags.StringVar(&registerReq.Type, RegisterTypeFlag, registerReq.Type, "The type of the user to register, should be user/app/peer")
	flags.StringVar(&registerReq.Affiliation, RegisterAffiliationFlag, registerReq.Affiliation, "The Org where the user affiliated")
	flags.IntVar(&registerReq.MaxEnrollments, RegisterMaxEnrollmentsFlag, registerReq.MaxEnrollments, "The max time for user to enroll, '0' means infinite times")
	//flags.StringArrayVar(&registerReq.Attributes, RegisterAttributesFlag, registerReq.Attributes, "The attributes which user want to use for fulture, it's a key-value arrays, Example {'k1':'v1','k2':'v2'}")
	return certRegisterCmd
}

type registerAction struct {
	common.ActionImpl
}

func newRegisterAction(flags *pflag.FlagSet) (*registerAction, error) {
	action := &registerAction{}
	err := action.Initialize(flags)
	return action, err
}

func (action *registerAction) register() error {
	fmt.Printf("Registering user [%s]\n", registerReq.Name)
	caClient, err := fabricCaClient.NewFabricCAClient()
	if err != nil {
		return err
	}
	registrar, err := action.Client().LoadUserFromStateStore("admin")
	if err != nil {
		return err
	}
	secret, err := caClient.Register(registrar, registerReq)
	if err != nil {
		return err
	} else {
		fmt.Printf("User register successfuly, Return: {Name: %s, EnrollmentSecret: %s}\n", registerReq.Name, secret)
	}

	return nil
}
