package user

import "testing"

func TestUser(t *testing.T) {
	login("hxy", "hxy", t)
}

func login(name, passwd string, t *testing.T) {
	u, err := Login(name, passwd)
	if err != nil {
		t.Fatal(err)
	}
	if u == nil {
		t.Fatal("Incorrect username or password")
	} else {
		t.Log(u)
	}
}
