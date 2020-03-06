package urpv

import (
	"crypto/rsa"
)

type Tokener interface {
	GetClaims() map[string]string
	HasClaim(k string) bool
	GetClaim(k string) string
	Encode(pubKey *rsa.PublicKey) ([]byte, error)
	Expired() bool
}

type Resourcer interface {
	GetName() string
	ValidateCaller(token Tokener) error
}

type Clienter interface {
	GetName() string
	GetOrigin() string
	GetCallback() string
	Validate(rs ResourceStore, t Tokener) (int, interface{})
}

type Storer interface {
	GetResourceStore() ResourceStore
	GetClientStore() ClientStore
}

type ResourceStore interface {
	GetResources(name ...string) []Resourcer
	//CreateResource(resc Resourcer) error
}

type ClientStore interface {
	GetClient(name, origin string) Clienter
}
