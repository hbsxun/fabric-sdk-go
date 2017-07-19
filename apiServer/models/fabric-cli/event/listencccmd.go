/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package event

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/api/apifabclient"
	"github.com/hyperledger/fabric-sdk-go/apiServer/models/fabric-cli/common"
	"github.com/spf13/pflag"
)

/*
var listenccCmd = &cobra.Command{
	Use:   "listencc",
	Short: "Listen to chaincode events.",
	Long:  "Listen to chaincode events",
	Run: func(cmd *cobra.Command, args []string) {
		if common.Config().ChaincodeID() == "" {
			fmt.Printf("\nMust specify the chaincode ID\n\n")
			cmd.HelpFunc()(cmd, args)
			return
		}
		if common.Config().ChaincodeEvent() == "" {
			fmt.Printf("\nMust specify the event name\n\n")
			cmd.HelpFunc()(cmd, args)
			return
		}

		action, err := newListenCCAction(cmd.Flags())
		if err != nil {
			fmt.Printf("\nError while initializing listenCCAction: %v\n", err)
			return
		}

		defer action.Terminate()

		err = action.invoke()
		if err != nil {
			fmt.Printf("\nError while running listenCCAction: %v\n", err)
		}
	},
}

func getListenCCCmd() *cobra.Command {
	flags := listenccCmd.Flags()
	common.Config().InitPeerURL(flags, "", "The URL of the peer on which to listen for events, e.g. localhost:7051")
	common.Config().InitChaincodeID(flags)
	common.Config().InitChaincodeEvent(flags)
	return listenccCmd
}
*/
type ListenCCArgs struct {
	PeerUrl        string `json:"peerUrl"`
	ChaincodeID    string `json:"chaincodeID"`
	ChaincodeEvent string `json:"chaincodeEvent"`
}

type listenccAction struct {
	common.Action
	inputEvent
}

func NewListenCCAction(args *ListenCCArgs) (*listenccAction, error) {
	flags := &pflag.FlagSet{}
	common.Config().InitPeerURL(flags, args.PeerUrl, "The URL of the peer on which to listen for events, e.g. localhost:7051")
	common.Config().InitChaincodeID(flags, args.ChaincodeID)
	common.Config().InitChaincodeEvent(flags, args.ChaincodeEvent)

	action := &listenccAction{inputEvent: inputEvent{done: make(chan bool)}}
	err := action.Initialize(flags)
	return action, err
}

func (action *listenccAction) Execute() error {
	eventHub, err := action.EventHub()
	if err != nil {
		return err
	}

	fmt.Printf("Registering CC event on chaincode [%s] and event [%s]\n", common.Config().ChaincodeID(), common.Config().ChaincodeEvent())

	registration := eventHub.RegisterChaincodeEvent(common.Config().ChaincodeID(), common.Config().ChaincodeEvent(), func(event *apifabclient.ChaincodeEvent) {
		fmt.Printf("Received CC event:\n")
		fmt.Printf("- Channel ID: %s\n", event.ChannelID)
		fmt.Printf("- Chaincode ID: %s\n", event.ChaincodeID)
		fmt.Printf("- Event: %s\n", event.EventName)
		fmt.Printf("- TxID: %s\n", event.TxID)
		fmt.Printf("- Payload: %v\n", event.Payload)

		fmt.Println("Press <enter> to terminate")
	})

	action.WaitForEnter()

	fmt.Printf("Unregistering CC event on chaincode [%s] and event [%s]\n", common.Config().ChaincodeID(), common.Config().ChaincodeEvent())
	eventHub.UnregisterChaincodeEvent(registration)

	return nil
}
