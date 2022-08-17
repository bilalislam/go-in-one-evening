package conditionals

var Password = "current-password"

func ResetPassword(code int) {
	if code == 0 || code == 1000 {
		return
	}

	Password = "new-password"
}
