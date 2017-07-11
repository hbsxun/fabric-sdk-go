package user

import "testing"

func TestUser(t *testing.T) {
	login("hxy", "hxy", t)
}

func login(name, passwd string, t *testing.T) {
	ok, err := Login(name, passwd)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("Incorrect username or password")
	}
}
