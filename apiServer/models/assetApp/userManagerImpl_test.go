package assetApp

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/hyperledger/fabric-sdk-go/apiServer/models/user"
)

func TestUserManager(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	name := "bill" + strconv.Itoa(r.Intn(1000)) //note: the len(name) in database is 18. len(passwd) is 20 bytes
	register(name, t)
	login(name, name, t) //passwd=name
	updateUser(name, t)
	getUserByName(name, t)
	getUserById(3, t)
}

func register(name string, t *testing.T) {
	impl := UserManagerImpl{}

	u := &user.User{
		Name:   name,
		Passwd: name,
		Email:  name + "@linux.com",
		Phone:  "123456",
	}

	id, ok := impl.Register(u)
	if !ok {
		t.Error("Register failed")
	} else {
		t.Log(id)
	}
}

func login(name, passwd string, t *testing.T) {
	impl := UserManagerImpl{}
	signedToken, err := impl.Login(&user.Secret{name, passwd})
	if err != nil {
		t.Fatal("Login failed", err)
	} else {
		t.Log("signedToken", signedToken)
	}
}

func updateUser(name string, t *testing.T) {
	impl := UserManagerImpl{}
	u := &user.User{
		Name:  name,
		Email: "billhan@linux.com",
		Phone: "110",
	}
	err := impl.UpdateInfo(u)
	if err != nil {
		t.Error("Update failed")
	} else {
		fmt.Println("update success")
	}
}

func getUserByName(name string, t *testing.T) {
	impl := UserManagerImpl{}
	userInfo, err := impl.GetUserInfoByName(name)
	if err != nil {
		t.Errorf("getUserByName failed")
	} else {
		t.Log("user: ", userInfo)
	}
}
func getUserById(id int, t *testing.T) {
	impl := UserManagerImpl{}
	userInfo, err := impl.GetUserInfoById(id)
	if err != nil {
		t.Errorf("getUserById failed")
	} else {
		t.Log("user: ", userInfo)
	}
}
