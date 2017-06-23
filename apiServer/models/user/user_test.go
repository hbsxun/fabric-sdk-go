package user

import (
	"log"
	"strconv"
	"testing"
)

func TestAddUser(*testing.T) {
	a := &User{
		Name:   "alice",
		Passwd: "alice",
		Mail:   "alice@example.com",
		Cert:   "jfwojsfj==",
	}
	id, _ := AddUser(a)
	log.Println(id)

	b, _ := GetUser(strconv.Itoa(id))

	log.Println(b)
}
