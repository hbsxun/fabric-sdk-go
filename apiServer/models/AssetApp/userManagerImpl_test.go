package AssetApp

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
)

func TestRegister(t *testing.T) {
	impl := UserManagerImpl{}

	u := &user.User{
		Name:   "bill",
		Passwd: "bill",
		Email:  "bill@linux.com",
		Phone:  "111",
	}

	id, ok := impl.Register(u)
	if !ok {
		t.Error("Register failed")
	} else {
		t.Log(id)
	}
}
