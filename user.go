package urpv

type user struct {
	Username string
	Password []byte
	Claims map[string]string
}

func (u user) GetUsername() string {
	return u.Username
}

func (u user) GetRequestedClaims(claims ...string) map[string]string {
	result := make(map[string]string)
	for _, c := range claims	 {
		if c == "email" {
			result["email"] = u.Username
		}

		result[c] = u.Claims[c]
	}

	return result
}

func (u user) IsVerified() bool {
	return true
}

func (u user) ConfirmPassword(password string) bool {
	return string(u.Password) == password
}