package urpv

import (
	"testing"
	"time"
)

// 1. Client attempts to call Resource without any id_code, must fail
func TestResourceCall_NoIDCode_MustFail(t *testing.T) {

	//info? POST token
}

// 2. Unknown Client attempts to obtain an id_code
func TestClientAuthorize_NoUser_MustPass(t *testing.T) {
	//authorize? GET
}

// 3. Client has been verified, obtains id_code
func TestResourceCall_HasIDCode_MustPass(t *testing.T) {
	//info? POST token
}

// 3b. Unknown Client attempts to obtain an id_code, and authenticate a user
func TestClientAuthorize_WithUser_MustPass(t *testing.T) {
	//authorize?
}

// Obtain Id Code (Authorize) -- Validate client, if requires 'user'-scope goto login. Generates idcode
// Login (If Required), adds requested user claims
// Call Resource with Token -- validates Token & claims

func simpleWebApp() Clienter {
	return NewClient("urpv_simple", "simple.localhost", []string{"client_info"})
}

func userWebApp() Clienter {
	return NewClient("urpv_user", "user.localhost", []string{"user_info"})
}

func getsimpleToken() Tokener {
	clms := make(map[string]string)
	clms["client_id"] = "urpv_simple"

	return token{
		Claims:    clms,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}
}

func getuserToken() Tokener {
	clms := make(map[string]string)
	clms["client_id"] = "urpv_user"
	clms["username"] = "userX"
	clms["email"] = "user@x.com"

	return token{
		Claims:    clms,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}
}

func TestAuthentication_Client_Simple(t *testing.T) {
	str := NewFakeStore()

	wa := simpleWebApp()
	err := wa.Validate(str.GetResourceStore(), getsimpleToken())

	if err != nil {
		t.Error(err)
		return
	}
}

func TestAuthentication_Client_User(t *testing.T) {
	str := NewFakeStore()

	wa := userWebApp()
	err := wa.Validate(str.GetResourceStore(), getuserToken())

	if err != nil {
		t.Error(err)
		return
	}
}
