package urpv

import (
	"errors"
	"fmt"
)

type resource struct {
	Name     string
	Requires []string //scopes required
}
func (r resource) GetName() string {
	return r.Name
}

//Call will validate whether the token is allowed to be used against the requested resource
func (r resource) ValidateCaller(token Tokener) error {
	if token.Expired() {
		return errors.New("token has expired")
	}

	for _, scp := range r.Requires {
		if !token.HasClaim(scp) {
			return errors.New(fmt.Sprintf("required scope %s not found", scp))
		}
	}

	return nil
}
