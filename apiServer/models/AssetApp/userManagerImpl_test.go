package AssetApp

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
)

func TestRegister(t *testing.T) {
	impl := UserManagerImpl{}

	u := &user.User{
		Name:   "lovecrypto04",
		Passwd: "dandan",
		Email:  "hh@linux.com",
		Phone:  "110",
	}

	id, ok := impl.Register(u)
	if !ok {
		t.Error("Register failed")
	} else {
		t.Log(id)
	}
}

/*func TestLogin(t *testing.T) {
	impl := UserManagerImpl{}
	ok := impl.Login("bill", "bill")
	if !ok {
		t.Error("Login failed")
	} else {
		fmt.Println("login success: bill")
	}
}

func TestUpdateinfo(t *testing.T) {
	impl := UserManagerImpl{}
	u := &user.User{
		Name:   "bill",
		Passwd: "bill",
		Email:  "billhan@linux.com",
		Phone:  "110",
	}
	err := impl.UpdateInfo(u)
	if err != nil {
		t.Error("Update failed")
	} else {
		fmt.Println("update success")
	}
}*/
