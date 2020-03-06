package urpv

import (
	"net/http"
)

type client struct {
	Name     string
	Resources []string
	Origin   string
}

func NewClient(name, origin string, resources []string) Clienter {
	return client{
		Name:     name,
		Resources: resources,
		Origin:   origin,
	}
}

func (c client) GetName() string {
	return c.Name
}

func (c client) GetOrigin() string {
	return c.Origin
}

func (c client) GetCallback() string {
	return c.Origin + "/signin"
}

func (c client) Validate(rs ResourceStore, t Tokener) (int, interface{}) {
	// fetch client by name
	// validate session endpoint matches
	// does audience require a user? ---> redirect
	//audience := rs.GetResources(c.Audience...)

	return http.StatusOK, nil
}
