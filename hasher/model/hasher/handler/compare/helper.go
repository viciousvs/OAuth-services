package compare

func validate(passwordHash, password string) bool {
	return len(passwordHash) < 8 || len(password) < 8
}
