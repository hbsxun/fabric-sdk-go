/*
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
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-ca/api"
	config "github.com/hyperledger/fabric-sdk-go/config"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	kvs "github.com/hyperledger/fabric-sdk-go/fabric-client/keyvaluestore"
	"github.com/hyperledger/fabric/bccsp"
	bccspFactory "github.com/hyperledger/fabric/bccsp/factory"

	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
)

type Member struct {
	caClient  fabricCAClient.Services
	adminUser fabricClient.User
}

func NewMember() *Member {
	admin := new(Member)
	err := admin.AdminEnroll("admin", "adminpw")
	if err != nil {
		log.Printf("Get admin Ecert error [%v]", err)
		return nil
	}
	return admin
}

// This test loads/enrols an admin user
// Using the admin, it registers, enrols, and revokes a test user
func (t *Member) AdminEnroll(adminName, adminPasswd string) error {
	testSetup := &BaseSetupImpl{
		//ConfigFile: "../fixtures/config/config_test.yaml",
		ConfigFile: "/home/hxy/gopath/src/github.com/hyperledger/fabric-sdk-go/test/fixtures/config/config_test.yaml",
	}
	testSetup.InitConfig()
	client := fabricClient.NewClient()

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
	if err != nil {
		return fmt.Errorf("Failed getting ephemeral software-based BCCSP [%s]", err)
	}

	cryptoSuite := bccspFactory.GetDefault()

	client.SetCryptoSuite(cryptoSuite)
	stateStore, err := kvs.CreateNewFileKeyValueStore("/tmp/enroll_user")
	if err != nil {
		return fmt.Errorf("CreateNewFileKeyValueStore return error[%s]", err)
	}
	client.SetStateStore(stateStore)

	t.caClient, err = fabricCAClient.NewFabricCAClient()
	if err != nil {
		return fmt.Errorf("NewFabricCAClient return error: %v", err)
	}

	// Admin user is used to register, enrol and revoke a test user
	t.adminUser, err = client.LoadUserFromStateStore("admin")

	if err != nil {
		return fmt.Errorf("client.GetUserContext return error: %v", err)
	}
	if t.adminUser == nil {
		key, cert, err := t.caClient.Enroll(adminName, adminPasswd)
		if err != nil {
			return fmt.Errorf("Enroll return error: %v", err)
		}
		if key == nil {
			return errors.New("private key return from Enroll is nil")
		}
		if cert == nil {
			return errors.New("cert return from Enroll is nil")
		}

		certPem, _ := pem.Decode(cert)
		if err != nil {
			return fmt.Errorf("pem Decode return error: %v", err)
		}

		cert509, err := x509.ParseCertificate(certPem.Bytes)
		if err != nil {
			return fmt.Errorf("x509 ParseCertificate return error: %v", err)
		}
		if cert509.Subject.CommonName != "admin" {
			return errors.New("CommonName in x509 cert is not the enrollmentID")
		}

		keyPem, _ := pem.Decode(key)
		if err != nil {
			return fmt.Errorf("pem Decode return error: %v", err)
		}
		t.adminUser = fabricClient.NewUser("admin")
		k, err := client.GetCryptoSuite().KeyImport(keyPem.Bytes, &bccsp.ECDSAPrivateKeyImportOpts{Temporary: false})
		if err != nil {
			return fmt.Errorf("KeyImport return error: %v", err)
		}
		t.adminUser.SetPrivateKey(k)
		t.adminUser.SetEnrollmentCertificate(cert)
		err = client.SaveUserToStateStore(t.adminUser, false)
		if err != nil {
			return fmt.Errorf("client.SetUserContext return error: %v", err)
		}
		t.adminUser, err = client.LoadUserFromStateStore("admin")
		if err != nil {
			return fmt.Errorf("client.GetUserContext return error: %v", err)
		}
		if t.adminUser == nil {
			return errors.New("client.GetUserContext return nil")
		}
	}
	return nil
}

//RegisterUser futural use will input a unique UserName
func (t *Member) RegisterUser(regReq *fabricCAClient.RegistrationRequest) (name, secret string, err error) {
	// Register a random user
	/*
		userName := t.createRandomName()
		registerRequest := fabricCAClient.RegistrationRequest{Name: userName, Type: "user", Affiliation: "org1.department1"}
	*/
	enrolmentSecret, err := t.caClient.Register(t.adminUser, regReq)
	if err != nil {
		return regReq.Name, "", fmt.Errorf("Register failed: %s", err)
	}
	//fmt.Printf("Registered User: %s, Secret: %s\n", userName, enrolmentSecret)

	return regReq.Name, enrolmentSecret, nil
}

func (t *Member) RevokeUser(revReq *fabricCAClient.RevocationRequest) error {
	/*
			revokeRequest := fabricCAClient.RevocationRequest{
				Name: userName,
			}
		return t.caClient.Revoke(t.adminUser, &revokeRequest)
	*/
	return t.caClient.Revoke(t.adminUser, revReq)
}

//UserEnroll in production this method should be invoked by the user, here for test
func (t *Member) UserEnroll(userName, enrolmentSecret string) (key, cert []byte, err error) {
	/*
		copClient, err = fabricCAClient.NewFabricCAClient()
		if err != nil {
			return fmt.Errorf("NewFabricCAClient return error: %v", err)
		}
	*/
	// Enrol the previously registered user
	key, cert, err = t.caClient.Enroll(userName, enrolmentSecret)
	if err != nil {
		return nil, nil, fmt.Errorf("Error enroling user: %s", err.Error())
	}
	//fmt.Printf("key [%s], cert [%s]", string(key), string(cert))
	return key, cert, nil

}
func (t *Member) UserEnrollWithCSR(req *api.EnrollmentRequest) (key, cert []byte, err error) {
	key, cert, err = t.caClient.EnrollWithCSR(req)
	if err != nil {
		return nil, nil, err
	}
	return key, cert, nil
}
func (t *Member) createRandomName() string {
	rand.Seed(time.Now().UnixNano())
	return "user" + strconv.Itoa(rand.Intn(500000))
}
