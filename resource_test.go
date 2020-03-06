package urpv

import (
	"testing"
	"time"
)

func TestValidateCaller_MustbeValid(t *testing.T) {
	prvKey, err := GenerateKeyPair()

	if err != nil {
		t.Fatal(err)
		return
	}

	clms := make(map[string]string)
	clms["client"] = "urpv_test"

	dumbTokn := token{
		Claims: clms,
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	raw, err := dumbTokn.Encode(&prvKey.PublicKey)

	if err != nil {
		t.Error(err)
		return
	}

	token, err := DecodeIDCode(raw, prvKey)

	if err != nil {
		t.Fatal(err)
		return
	}

	resrc := resource{
		Name:     "urpv_resource",
		Requires: []string{
			"client",
		},
	}

	err = resrc.ValidateCaller(token)

	if err != nil {
		t.Fatal(err)
		return
	}
}
