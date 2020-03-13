package urpv

type fakeStore struct {
	Clients   ClientStore
	Resources ResourceStore
	Users UserStore
}

func NewFakeStore() Storer {
	return fakeStore{
		Clients:   newFakeClientStore(),
		Resources: newFakeResourceStore(),
		Users: newFakeUserStore(),
	}
}

func (s fakeStore) GetClientStore() ClientStore {
	return s.Clients
}

func (s fakeStore) GetResourceStore() ResourceStore {
	return s.Resources
}

func (s fakeStore) GetUserStore() UserStore {
	return s.Users
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

func (cs fakeClientStore) GetClient(name string) Clienter {
	for _, v := range cs.clients {
		if v.GetName() == name {
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

type fakeUserStore struct {
	users map[string]Userer
}

func newFakeUserStore() fakeUserStore{
	users := map[string]Userer{
		"0X":user{
			Username: "userX",
			Password: []byte("passwordX"),
			Claims: map[string]string{
				"email": "user@x.com",
				"nickname": "User, X",
			},
		},
		"1N":user{
			Username: "userN",
			Password: []byte("passwordN"),
			Claims: map[string]string{
				"email": "user@n.com",
				"nickname": "User, N",
			},
		},
	}
	return fakeUserStore{users: users}
}

func (us fakeUserStore) Login(username, password string) Userer {
	key := "!"

	for k, v := range us.users {
		if v.GetUsername() == username {
			key = k
			break
		}
	}

	if key == "!" {
		return nil
	}

	usr, ok := us.users[key]

	if !ok {
		return nil
	}

	if !usr.ConfirmPassword(password) {
		return nil
	}

	return usr
}
func (us fakeUserStore) GetClaimValue(userKey string, claims ...string) map[string]string {
	usr, ok := us.users[userKey]

	if !ok {
		return nil
	}

	return usr.GetRequestedClaims(claims...)
}