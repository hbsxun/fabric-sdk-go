package models

import (
	"encoding/base64"

	"github.com/hyperledger/fabric-ca/api"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
)

var admin *sdkIgn.Member

type RegisterResponse struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

type EnrollResponse struct {
	Key  string `json:"key"`  //base64
	Cert string `json:"cert"` //base64
}

//Enroll returns key, cert, err
func Enroll(req *api.EnrollmentRequest) (*EnrollResponse, error) {
	key, cert, err := admin.UserEnrollWithCSR(req)
	keyStr := base64.StdEncoding.EncodeToString(key)
	certStr := base64.StdEncoding.EncodeToString(cert)
	if err != nil {
		return nil, err
	}
	return &EnrollResponse{keyStr, certStr}, nil
}

//Register returns name, secret and error
func Register(req *fabricCAClient.RegistrationRequest) (*RegisterResponse, error) {
	name, secret, err := admin.RegisterUser(req)
	if err != nil {
		return nil, err
	}
	return &RegisterResponse{name, secret}, nil
}

func init() {
	admin = sdkIgn.NewMember()
}
