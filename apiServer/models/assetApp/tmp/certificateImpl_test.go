package AssetApp

import (
	"fmt"
	"testing"
)

func TestGetIdentityandSaveToDB(t *testing.T) {
	impl := CertificateImpl{}
	key, cert := impl.GetIdentity("warm")
	if key == "" && cert == "" {
		t.Error("GetIdentity failed")
	} else {
		fmt.Printf("key: %s, cert: %s\n", key, cert)
		//err := impl.SaveToDB(key, cert, "lovecrypto04")
	}
}
