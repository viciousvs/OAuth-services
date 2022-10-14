package generate

import (
	"github.com/viciousvs/OAuth-services/hasher/handler"
	"github.com/viciousvs/OAuth-services/hasher/utils/customErors"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
	"golang.org/x/crypto/bcrypt"
)

func Handle(req *hasherPb.Password) (string, error) {
	if req == nil {
		return "", customErors.ErrNilRequest
	}
	password := req.GetValue()
	if len(password) < 8 {
		return "", customErors.ErrInvalidData
	}

	return handler.NewBcryptHasher(bcrypt.DefaultCost).GenerateHash(password)
}
