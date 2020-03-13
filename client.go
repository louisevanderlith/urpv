package urpv

type client struct {
	Name      string
	Resources []string
	Origin    string
}

func NewClient(name, origin string, resources []string) Clienter {
	return client{
		Name:      name,
		Resources: resources,
		Origin:    origin,
	}
}

func (c client) GetName() string {
	return c.Name
}

func (c client) GetOrigin() string {
	return c.Origin
}

func (c client) GetResources() []string {
	return c.Resources
}

func (c client) GetCallback() string {
	return c.Origin + "/signin"
}

func (c client) Validate(rs ResourceStore, t Tokener) error {
	rr := rs.GetResources(c.GetResources()...)

	for _, rsc := range rr {
		err := rsc.ValidateCaller(t)

		if err != nil {
			return err
		}
	}

	return nil
}
