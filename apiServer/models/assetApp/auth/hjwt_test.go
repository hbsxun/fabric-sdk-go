package auth

import "testing"

func TestJWT(t *testing.T) {
	signedToken := CreateToken(0, "hxy", true)
	if signedToken == "" {
		t.Fatal("create token failed")
	}

	if valid := IsTokenValid(signedToken); !valid {
		t.Fatal("Token is invalid")
	}

	if isAdmin := IsAdmin(signedToken); !isAdmin {
		t.Fatal("User should be an admin")
	}

	//t.Logf("SignedToken: %s\t Valid: %v\t | isAdmin: %v\t isEnrolled: %v\n", signedToken, valid, isAdmin, isEnrolled)

}
