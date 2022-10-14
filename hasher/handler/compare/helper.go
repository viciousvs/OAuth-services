package compare

func validate(passwordHash, password string) bool {
	hashLen, pwLen := len(passwordHash), len(password)
	return hashLen > pwLen && pwLen >= 8
}
