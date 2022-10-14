package handler

import (
	"golang.org/x/crypto/bcrypt"
)

type bcryptHasher struct {
	cost int
}

func NewBcryptHasher(cost int) bcryptHasher {
	return bcryptHasher{cost: cost}
}
func (b bcryptHasher) GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (b bcryptHasher) CompareHashAndPassword(passwordHash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
}
