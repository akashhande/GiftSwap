package security

func IsValidCredentials(username, password string) bool {
	if username == "root" && password == "root" {
		return true
	}
	return false
}
