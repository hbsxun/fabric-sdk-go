package user

import (
	"log"
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

	b, _ := GetUser(id)

	log.Println(b)
}
