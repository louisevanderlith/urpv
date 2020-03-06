package urpv

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"time"
)

type token struct {
	Claims map[string]string `json:"claims"`
	ExpiresAt time.Time `json:"expires_at"`
}

func DecodeIDCode(idcode []byte, prvKey *rsa.PrivateKey) (Tokener, error) {
	dcryptd, err := rsa.DecryptPKCS1v15(rand.Reader, prvKey, idcode)

	if err != nil {
		return nil, err
	}

	result := token{}
	err = json.Unmarshal(dcryptd, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t token) GetClaims() map[string]string {
	return t.Claims
}

func (t token) HasClaim(k string) bool {
	_, ok  := t.Claims[k]

	return ok
}

func (t token) GetClaim(k string) string {
	return t.Claims[k]
}

func (t token) Encode(pubKey *rsa.PublicKey) ([]byte, error) {
	bits, err := json.Marshal(t)

	if err != nil {
		return nil, err
	}

	return rsa.EncryptPKCS1v15(rand.Reader, pubKey, bits)
}

func (t token) Expired() bool {
	return t.ExpiresAt.Before(time.Now())
}