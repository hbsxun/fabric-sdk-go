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
	"os"
	"path"

	"github.com/hyperledger/fabric-sdk-go/config"
	"github.com/hyperledger/fabric-sdk-go/fabric-client/events"
	fcUtil "github.com/hyperledger/fabric-sdk-go/fabric-client/helpers"

	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	bccspFactory "github.com/hyperledger/fabric/bccsp/factory"
)

// BaseSetupImpl implementation of BaseTestSetup
type BaseSetupImpl struct {
	Client          fabricClient.Client
	Chain           fabricClient.Chain
	EventHub        events.EventHub
	ConnectEventHub bool
	ConfigFile      string
	ChainID         string
	ChainCodeID     string
	Initialized     bool
	ChannelConfig   string
}

func NewBaseSetupImpl(prefix string) *BaseSetupImpl {
	testSetup := &BaseSetupImpl{
		ConfigFile:      prefix + "/fixtures/config/config_test.yaml",
		ChainID:         "testchannel",
		ChannelConfig:   prefix + "/fixtures/channel/testchannel.tx",
		ConnectEventHub: true,
	}

	if err := testSetup.Initialize(); err != nil {
		panic("BaseSetupImpl Initialize failed...")
		return nil
	}
	return testSetup
}

// Initialize reads configuration from file and sets up client, chain and event hub
func (setup *BaseSetupImpl) Initialize() error {

	if err := setup.InitConfig(); err != nil {
		return fmt.Errorf("Init from config failed: %v", err)
	}

	// Initialize bccsp factories before calling get client
	/*
		err := bccspFactory.InitFactories(&bccspFactory.FactoryOpts{
			ProviderName: "SW",
			SwOpts: &bccspFactory.SwOpts{
				HashFamily: config.GetSecurityAlgorithm(),
				SecLevel:   config.GetSecurityLevel(),
				FileKeystore: &bccspFactory.FileKeystoreOpts{
					KeyStorePath: config.GetKeyStorePath(),
				},
				Ephemeral: false,
			},
		})
	*/
	err := bccspFactory.InitFactories(&bccspFactory.FactoryOpts{
		ProviderName: "PKCS11",
		Pkcs11Opts: &bccspFactory.PKCS11Opts{
			HashFamily: config.GetSecurityAlgorithm(),
			SecLevel:   config.GetSecurityLevel(),
			FileKeystore: &bccspFactory.FileKeystoreOpts{
				KeyStorePath: config.GetKeyStorePath(),
			},
			Ephemeral: false,
			Library:   "/usr/lib/softhsm/libsofthsm2.so",
			Pin:       "1234",
			Label:     "forfabric",
		},
	})

	if err != nil {
		return fmt.Errorf("Failed getting ephemeral software-based BCCSP [%s]", err)
	}

	client, err := fcUtil.GetClient("admin", "adminpw", "/tmp/enroll_user")
	if err != nil {
		return fmt.Errorf("Create client failed: %v", err)
	}
	setup.Client = client

	chain, err := fcUtil.GetChain(setup.Client, setup.ChainID)
	if err != nil {
		return fmt.Errorf("Create chain (%s) failed: %v", setup.ChainID, err)
	}
	setup.Chain = chain

	// Create and join channel
	if err := fcUtil.CreateAndJoinChannel(client, chain, setup.ChannelConfig); err != nil {
		return fmt.Errorf("CreateAndJoinChannel return error: %v", err)
	}

	eventHub, err := getEventHub()
	if err != nil {
		return err
	}

	if setup.ConnectEventHub {
		if err := eventHub.Connect(); err != nil {
			return fmt.Errorf("Failed eventHub.Connect() [%s]", err)
		}
	}
	setup.EventHub = eventHub

	setup.Initialized = true

	return nil
}

// InitConfig ...
func (setup *BaseSetupImpl) InitConfig() error {
	if err := config.InitConfig(setup.ConfigFile); err != nil {
		return err
	}
	return nil
}

// InstantiateCC ...
func (setup *BaseSetupImpl) InstantiateCC(chainCodeID string, chainID string, chainCodePath string, chainCodeVersion string, args []string) error {
	if err := fcUtil.SendInstantiateCC(setup.Chain, chainCodeID, chainID, args, chainCodePath, chainCodeVersion, []fabricClient.Peer{setup.Chain.GetPrimaryPeer()}, setup.EventHub); err != nil {
		return err
	}
	return nil
}

// InstallCC ...
func (setup *BaseSetupImpl) InstallCC(chainCodeID string, chainCodePath string, chainCodeVersion string, chaincodePackage []byte) error {
	if err := fcUtil.SendInstallCC(setup.Client, setup.Chain, chainCodeID, chainCodePath, chainCodeVersion, chaincodePackage, setup.Chain.GetPeers(), setup.GetDeployPath()); err != nil {
		return fmt.Errorf("SendInstallProposal return error: %v", err)
	}
	return nil
}

// GetDeployPath ..
func (setup *BaseSetupImpl) GetDeployPath() string {
	pwd, _ := os.Getwd()
	return path.Join(pwd, "../fixtures")
}

// Query ...
func (setup *BaseSetupImpl) Query(chainID string, chainCodeID string, args []string) (string, error) {
	transactionProposalResponses, _, err := fcUtil.CreateAndSendTransactionProposal(setup.Chain, chainCodeID, chainID, args, []fabricClient.Peer{setup.Chain.GetPrimaryPeer()}, nil)
	if err != nil {
		return "", fmt.Errorf("CreateAndSendTransactionProposal return error: %v", err)
	}
	return string(transactionProposalResponses[0].GetResponsePayload()), nil
}

// getEventHub initilizes the event hub
func getEventHub() (events.EventHub, error) {
	eventHub := events.NewEventHub()
	foundEventHub := false
	peerConfig, err := config.GetPeersConfig()
	if err != nil {
		return nil, fmt.Errorf("Error reading peer config: %v", err)
	}
	for _, p := range peerConfig {
		if p.EventHost != "" && p.EventPort != 0 {
			fmt.Printf("******* EventHub connect to peer (%s:%d) *******\n", p.EventHost, p.EventPort)
			eventHub.SetPeerAddr(fmt.Sprintf("%s:%d", p.EventHost, p.EventPort),
				p.TLS.Certificate, p.TLS.ServerHostOverride)
			foundEventHub = true
			break
		}
	}

	if !foundEventHub {
		return nil, fmt.Errorf("No EventHub configuration found")
	}

	return eventHub, nil
}
