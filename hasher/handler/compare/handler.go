package compare

import (
	"github.com/viciousvs/OAuth-services/hasher/handler"
	"github.com/viciousvs/OAuth-services/hasher/utils/customErors"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
	"golang.org/x/crypto/bcrypt"
)

func Handle(req *hasherPb.CompareRequest) error {
	if req == nil {
		return customErors.ErrNilRequest
	}
	hash, psw := req.GetPasswordHash(), req.GetPassword()
	if !validate(hash, psw) {
		return customErors.ErrInvalidData
	}

	err := handler.NewBcryptHasher(bcrypt.DefaultCost).CompareHashAndPassword(hash, psw)
	if err != nil {
		return err
	}

	return nil
}
