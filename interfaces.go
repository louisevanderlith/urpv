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
	Validate(rs ResourceStore, t Tokener) error
}

type Userer interface {
	GetUsername() string
	IsVerified() bool
	ConfirmPassword(password string) bool
	GetRequestedClaims(claims ...string) map[string]string
}

type Storer interface {
	GetResourceStore() ResourceStore
	GetClientStore() ClientStore
	GetUserStore() UserStore
}

type ResourceStore interface {
	GetResources(name ...string) []Resourcer
}

type ClientStore interface {
	GetClient(name string) Clienter
}

type UserStore interface {
	Login(username, password string) Userer
	GetClaimValue(userKey string, claims ...string) map[string]string
}