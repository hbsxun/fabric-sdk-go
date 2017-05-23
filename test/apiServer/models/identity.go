package models

import (
	"encoding/base64"

	"github.com/hyperledger/fabric-ca/api"
	fabricCAClient "github.com/hyperledger/fabric-sdk-go/fabric-ca-client"
	sdkIgn "github.com/hyperledger/fabric-sdk-go/test/integration"
)

var admin *sdkIgn.Member

//Register Object
type RegisterRequest struct {
	Name           string      `json:"name"`
	Type           string      `json:"type"`
	MaxEnrollments int         `json:"maxEnrollments"`
	Attributes     []Attribute `json:"attributes"`
	Affiliation    string      `json:"affiliation"`
}
type Attribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type RegisterResponse struct {
	Name   string `json:"name"`
	Secret string `json:"secret"`
}

//Enroll Object
type EnrollRequest struct {
	Name   string  `json:"name"`
	Secret string  `json:"secret"`
	CSR    CSRInfo `json:"csr"`
}
type CSRInfo struct {
	CN    string   `json:"cn"`
	Hosts []string `json:"hosts"`
	//KeyRequest BasicKeyRequest `json:"keyRequest"`
}
type EnrollResponse struct {
	Key  string `json:"key"`  //base64
	Cert string `json:"cert"` //base64
}

//Enroll returns key, cert, err
func Enroll(req *EnrollRequest) (*EnrollResponse, error) {
	fbReq := &api.EnrollmentRequest{
		Name:   req.Name,
		Secret: req.Secret,
		CSR: &api.CSRInfo{
			CN:    req.CSR.CN,
			Hosts: req.CSR.Hosts,
		},
	}
	key, cert, err := admin.UserEnrollWithCSR(fbReq)
	keyStr := base64.StdEncoding.EncodeToString(key)
	certStr := base64.StdEncoding.EncodeToString(cert)
	if err != nil {
		return nil, err
	}
	return &EnrollResponse{keyStr, certStr}, nil
}

//Register returns name, secret and error
func Register(req *RegisterRequest) (*RegisterResponse, error) {
	var attrs []fabricCAClient.Attribute
	for _, ele := range req.Attributes {
		attr := fabricCAClient.Attribute{
			Key:   ele.Key,
			Value: ele.Value,
		}
		attrs = append(attrs, attr)
	}
	fbReq := &fabricCAClient.RegistrationRequest{
		Name:           req.Name,
		Type:           req.Type,
		MaxEnrollments: req.MaxEnrollments,
		Attributes:     attrs,
		Affiliation:    req.Affiliation,
	}
	name, secret, err := admin.RegisterUser(fbReq)
	if err != nil {
		return nil, err
	}
	return &RegisterResponse{name, secret}, nil
}
