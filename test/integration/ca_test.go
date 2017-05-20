package integration

import (
	"testing"

	"github.com/hyperledger/fabric-ca/api"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
)

var prefix = "/home/hxy/gopath/src/github.com/hyperledger/fabric-sdk-go/test"

func Test_RegisterUser(t *testing.T) {

	testSetup := NewBaseSetupImpl(prefix)
	admin := NewMember(testSetup.Client)

	reqReq := &fabricCAClient.RegistrationRequest{
		Name:           "billhxy",
		Type:           "user", //"user, app, peer"
		MaxEnrollments: 0,
		Affiliation:    "org1.department1",
		Attributes: []fabricCAClient.Attribute{
			fabricCAClient.Attribute{"Job", "Software Enginner"},
			fabricCAClient.Attribute{"Title", "Junior"},
		},
	}
	name, secret, err := admin.RegisterUser(reqReq)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		t.Logf("name [%s], secret [%s]", name, secret)
	}

	key, cert, err := admin.UserEnroll(name, secret)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		t.Logf("key: ", string(key))
		t.Logf("cert: ", string(cert))
	}

	enrolReq := &api.EnrollmentRequest{
		Name:   name,
		Secret: secret,
	}
	key2, cert2, err := admin.UserEnrollWithCSR(enrolReq)
	if err != nil {
		t.Fatal(err.Error())
	} else {
		t.Logf("key: ", string(key2))
		t.Logf("cert: ", string(cert2))
	}
}
