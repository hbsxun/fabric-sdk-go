package auth

import (
	"encoding/base64"
	"encoding/json"
)

type Identity struct {
	EnrollName   string `json:"enrollName"`
	EnrollSecret string `json:"enrollSecret"`
}

func Serialize(id Identity) string {
	idJsonByte, _ := json.Marshal(id)
	return base64.StdEncoding.EncodeToString(idJsonByte)
}

func UnSerialize(id string) (Identity, error) {
	var idStruct Identity
	idJsonByte, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		return Identity{}, err
	}
	err = json.Unmarshal(idJsonByte, &idStruct)
	if err != nil {
		return Identity{}, err
	}
	return idStruct, nil
}

func (this *Identity) GetEnrollName() string {
	return this.EnrollName
}

func (this *Identity) GetEnrollSecret() string {
	return this.EnrollSecret
}
