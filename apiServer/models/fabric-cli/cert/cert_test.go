package cert

import (
	"fmt"
	"testing"
)

func TestCert(t *testing.T) {
	username := "bob"
	secret := register(&RegisterArgs{
		Name:        username,
		Type:        "user",
		Affiliation: "org1.department1",
	}, t)

	enroll(&EnrollArgs{
		Name:   username,
		Secret: secret,
	}, t)
}

func register(args *RegisterArgs, t *testing.T) (secret string) {
	action, err := NewRegisterAction(args)
	if err != nil {
		t.Fatal(err)
	}

	secret, err = action.Execute()
	if err != nil {
		t.Fatal(err)
	}
	return secret
}

func enroll(args *EnrollArgs, t *testing.T) {
	action, err := NewEnrollAction(args)
	if err != nil {
		t.Fatal(err)
	}

	keyCert, err := action.Execute()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("key.cert: ", keyCert)
}
