package hjwt

import "testing"

func TestJWT(t *testing.T) {
	id, name, isAdmin := "0", "hxy", false
	signedToken := CreateToken(id, name, isAdmin)
	if signedToken == "" {
		t.Fatal("create token failed")
	}

	valid, isAdmin := CheckToken(signedToken)
	if !valid {
		t.Fatal("CheckToken failed")
	}
	if isAdmin {
		t.Fatal("should not be a Admin")
	}

	t.Logf("SignedToken: %s | Valid: %v | isAdmin: %v\n", signedToken, valid, isAdmin)

}
