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
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	EnrollNameFlag   = "Name"
	EnrollSecretFlag = "Secret"
)

var enrollReq *api.EnrollmentRequest

var certEnrollCmd = &cobra.Command{
	Use:   "enroll",
	Short: "User Enroll",
	Long:  "User Enroll for key and cert",
	Run: func(cmd *cobra.Command, args []string) {
		if enrollReq.Name == "" {
			fmt.Printf("\nMust specify the Enroll Name\n\n")
			cmd.HelpFunc()(cmd, args)
			return
		}
		if enrollReq.Secret == "" {
			fmt.Printf("\nMust specify the Enroll Secret\n\n")
			cmd.HelpFunc()(cmd, args)
			return
		}
		action, err := newEnrollAction(cmd.Flags())
		if err != nil {
			common.Logger.Criticalf("Error while initializing enrollAction: %v", err)
			return
		}
		err = action.enroll()
		if err != nil {
			common.Logger.Criticalf("Error while running enrollAction: %v", err)
			return
		}
	},
}

func getCertEnrollCmd() *cobra.Command {
	enrollReq = &api.EnrollmentRequest{}
	flags := certEnrollCmd.Flags()
	flags.StringVar(&enrollReq.Name, EnrollNameFlag, enrollReq.Name, "The user name to enroll")
	flags.StringVar(&enrollReq.Secret, EnrollSecretFlag, enrollReq.Secret, "The user secret of the user, which is returned from registration")
	return certEnrollCmd
}

type enrollAction struct {
	common.ActionImpl
}

func newEnrollAction(flags *pflag.FlagSet) (*enrollAction, error) {
	action := &enrollAction{}
	err := action.Initialize(flags)
	return action, err
}

func (action *enrollAction) enroll() error {
	fmt.Printf("enrolling user [%s]\n", registerReq.Name)
	caClient, err := fabricCaClient.NewFabricCAClient()
	if err != nil {
		return err
	}
	key, cert, err := caClient.Enroll(enrollReq.Name, enrollReq.Secret)
	if err != nil {
		return err
	} else {
		fmt.Printf("User enroll successfuly, Return: {Key: %s, Cert: %s}\n", base64.StdEncoding.EncodeToString(key), base64.StdEncoding.EncodeToString(cert))
	}

	return nil
}
