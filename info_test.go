package urpv

import (
	"crypto/rsa"
	"testing"
	"time"
)

func getExpiredToken(pubK *rsa.PublicKey) []byte {
	clms := make(map[string]string)
	clms["client_id"] = "urpv_simple"

	tkn := token{
		Claims:    clms,
		ExpiresAt: time.Now().Add(-2 * time.Hour),
	}

	k, err := tkn.Encode(pubK)

	if err != nil {
		panic(err)
	}

	return k
}

func getInvalidToken(pubK *rsa.PublicKey) []byte {
	clms := make(map[string]string)
	clms["client_id"] = "urpv_user"

	tkn := token{
		Claims:    clms,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	k, err := tkn.Encode(pubK)

	if err != nil {
		panic(err)
	}

	return k
}

func getValidToken(pubK *rsa.PublicKey) []byte {
	clms := make(map[string]string)
	clms["client_id"] = "urpv_user"
	clms["username"] = "userX"
	clms["email"] = "user@x.com"

	tkn := token{
		Claims:    clms,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}

	k, err := tkn.Encode(pubK)

	if err != nil {
		panic(err)
	}

	return k
}

func TestIntrospect_ExpiredToken_MustBe_UnAuthorized(t *testing.T) {
	kys, err := GenerateKeyPair()

	if err != nil {
		t.Error(err)
		return
	}

	tkn := getExpiredToken(&kys.PublicKey)
	_, err = Introspect(tkn, kys, NewFakeStore(), "client_info")

	if err == nil {
		t.Error("expected error")
		return
	}

	if err.Error() != "token has expired" {
		t.Error("unexpected error", err.Error())
		return
	}
}

func TestIntrospect_NonValidClaims_MustBe_UnAuthorized(t *testing.T) {
	kys, err := GenerateKeyPair()

	if err != nil {
		t.Error(err)
		return
	}

	tkn := getInvalidToken(&kys.PublicKey)
	_, err = Introspect(tkn, kys, NewFakeStore(), "user_info")

	if err == nil {
		t.Error("expected error")
		return
	}

	if err.Error() != "required scope username not found" {
		t.Error("unexpected error", err.Error())
		return
	}
}

func TestIntrospect_ValidToken_MustBe_Authorized(t *testing.T) {
	kys, err := GenerateKeyPair()

	if err != nil {
		t.Error(err)
		return
	}

	tkn := getValidToken(&kys.PublicKey)
	claims, err := Introspect(tkn, kys, NewFakeStore(), "user_info")

	if err != nil {
		t.Error(err)
		return
	}

	if claims["username"] == "" {
		t.Error("username claim not found")
	}
}
