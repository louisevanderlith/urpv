package urpv

import (
	"crypto/rsa"
	"github.com/pkg/errors"
	"log"
)

//Introspect returns all 'granted' claims contained within the access_code
func Introspect(raw []byte, privK *rsa.PrivateKey, s Storer, resources ...string) (map[string]string, error) {
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

	rsrc := s.GetResourceStore().GetResources(resources...)

	if len(rsrc) == 0 {
		log.Println(resources, "not found")
		return nil, errors.New("invalid resource")
	}

	clntId := toknr.GetClaim("client_id")
	clnt := s.GetClientStore().GetClient(clntId, )

	if clnt == nil {
		return nil, errors.New("invalid client")
	}

	err = clnt.Validate(s.GetResourceStore(), toknr)

	if err != nil {
		return nil, err
	}

	return toknr.GetClaims(), nil
}