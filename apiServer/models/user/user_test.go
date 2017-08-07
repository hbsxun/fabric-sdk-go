package user

import "testing"

func TestUser(t *testing.T) {
	login(&Secret{
		"hxy",
		"hxy",
	}, t)
}

func login(ss *Secret, t *testing.T) {
	u, err := Login(ss)
	if err != nil {
		t.Fatal(err)
	}
	if u == nil {
		t.Fatal("Incorrect username or password")
	} else {
		t.Log(u)
	}
}
