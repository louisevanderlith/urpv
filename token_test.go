package urpv

import (
	"testing"
)

func TestDecodeIDCode_HasClaims(t *testing.T) {
	prvKey, err := GenerateKeyPair()

	if err != nil {
		t.Fatal(err)
		return
	}

	clms := make(map[string]string)
	clms["client"] = "urpv_test"

	dumbTokn := token{
		Claims: clms,
	}

	raw, err := dumbTokn.Encode(&prvKey.PublicKey)

	if err != nil {
		t.Error(err)
		return
	}

	tkn, err := DecodeIDCode(raw, prvKey)

	if err != nil {
		t.Fatal(err)
		return
	}

	if !tkn.HasClaim("client") {
		t.Error("doesn't have correct claim")
		return
	}

	if tkn.GetClaim("client") != "urpv_test" {
		t.Errorf("unexpected client claim %s, expected urpv_test", tkn.GetClaim("client"))
	}
}

func TestTokenEncode(t *testing.T) {
	prvKey, err := GenerateKeyPair()

	if err != nil {
		t.Fatal(err)
		return
	}

	clms := make(map[string]string)
	clms["client"] = "urpv_test"

	dumbTokn := token{
		Claims: clms,
	}

	raw, err := dumbTokn.Encode(&prvKey.PublicKey)

	if err != nil {
		t.Error(err)
		return
	}

	if len(raw) == 0 {
		t.Error("encoded token is empty")
	}
}