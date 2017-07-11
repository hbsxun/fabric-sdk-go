package AssetApp

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
	ok := impl.Login(name, passwd)
	if !ok {
		t.Error("Login failed")
	} else {
		fmt.Println("login success: bill")
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
