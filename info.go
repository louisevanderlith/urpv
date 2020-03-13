package urpv

import (
	"crypto/rsa"
	"github.com/pkg/errors"
)

//Introspect returns all 'granted' claims contained within the access_code
func Introspect(origin string, raw []byte, privK *rsa.PrivateKey, s Storer) (map[string]string, error) {
	if len(raw) == 0 {
		return nil, errors.New("no token")
	}

	toknr, err := DecodeIDCode(raw, privK)

	if err != nil {
		return nil, err
	}

	if !toknr.HasClaim("client_id") {
		return nil, errors.New("invalid client_id")
	}

	clnt := s.GetClientStore().GetClient(toknr.GetClaim("client_id"), origin)

	if clnt == nil {
		return nil, errors.New("invalid client")
	}

	err = clnt.Validate(s.GetResourceStore(), toknr)

	if err != nil {
		return nil, err
	}

	return toknr.GetClaims(), nil
}