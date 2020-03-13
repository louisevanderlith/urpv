package urpv

type fakeStore struct {
	Clients   ClientStore
	Resources ResourceStore
}

func NewFakeStore() Storer {
	return fakeStore{
		Clients:   newFakeClientStore(),
		Resources: newFakeResourceStore(),
	}
}

func (s fakeStore) GetClientStore() ClientStore {
	return s.Clients
}

func (s fakeStore) GetResourceStore() ResourceStore {
	return s.Resources
}

type fakeClientStore struct {
	clients []Clienter
}

func newFakeClientStore() ClientStore {
	clients := []Clienter{
		NewClient("urpv_simple", "simple.localhost", []string{"client_info"}),
		NewClient("urpv_user", "user.localhost", []string{"user_info"}),
	}

	return fakeClientStore{clients: clients}
}

func (cs fakeClientStore) GetClient(name, origin string) Clienter {
	for _, v := range cs.clients {
		if v.GetName() == name && v.GetOrigin() == origin {
			return v
		}
	}

	return nil
}

type fakeResourceStore struct {
	resources []Resourcer
}

func newFakeResourceStore() ResourceStore {
	resources := []Resourcer{
		resource{
			Name: "client_info",
			Requires: []string{
				"client_id",
			},
		},
		resource{
			Name: "user_info",
			Requires: []string{
				"username",
				"email",
			},
		},
	}
	return fakeResourceStore{resources: resources}
}

func (rs fakeResourceStore) GetResources(name ...string) []Resourcer {
	var result []Resourcer

	for _, n := range name {
		for _, v := range rs.resources {
			if v.GetName() == n {
				result = append(result, v)
				break
			}
		}
	}

	return result
}
