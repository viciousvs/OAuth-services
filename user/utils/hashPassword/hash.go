package hashPassword

import (
	"github.com/viciousvs/OAuth-services/user/utils/customErors"
	"golang.org/x/crypto/bcrypt"
)

//Hash
func Hash(password string) (string, error) {
	//TODO Implement Password Hashing
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", customErors.ErrCouldntBeHashed
	}
	return string(hash), nil
}

//CompareHashes
func CompareHashAndPassword(hash, plain string) error {
	//TODO implement comapre
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain))
}
