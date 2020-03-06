package urpv

import (
	"net/http"
	"testing"
)

// 1. Client attempts to call Resource without any id_code, must fail
func TestResourceCall_NoIDCode_MustFail(t *testing.T) {

	//info? POST token
}

// 2. Unknown Client attempts to obtain an id_code
func TestClientAuthorize_NoUser_MustPass(t *testing.T){
	//authorize? GET
}

// 3. Client has been verified, obtains id_code
func TestResourceCall_HasIDCode_MustPass(t *testing.T){
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
	return NewClient("urpv_simple", "simple.localhost",  []string{"client_info"})
}

func userWebApp() Clienter {
	return NewClient("urpv_user", "user.localhost", []string{"user_info"})
}

func TestAuthentication_Client(t *testing.T) {
	str := NewFakeStore()

	wa :=  simpleWebApp()
	code, result := wa.Validate(str.GetResourceStore(), )

	if code != http.StatusOK {
		t.Error(result)
		return
	}

	t.Log(result)

	if result == nil {
		t.Error("result is empty")
	}
}
