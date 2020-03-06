package urpv

import (
	"crypto/rand"
	"crypto/rsa"
)

//GenerateKeyPair creates a new Pub & Private key pair to be used for signing
func GenerateKeyPair() (*rsa.PrivateKey, error){
	reader := rand.Reader

	privKey, err := rsa.GenerateKey(reader, 2048)

	if err != nil {
		return nil, err
	}

	err = privKey.Validate()
	if err != nil {
		return nil, err
	}

	return privKey, nil
}


