package compare

import (
	"github.com/viciousvs/OAuth-services/hasher/service/argon2"
	"github.com/viciousvs/OAuth-services/hasher/utils/customErors"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
)

func Handle(req *hasherPb.CompareRequest) (bool, error) {
	if req == nil {
		return false, customErors.ErrNilRequest
	}
	hash, psw := req.GetPasswordHash(), req.GetPassword()
	if !validate(hash, psw) {
		return false, customErors.ErrInvalidData
	}

	return argon2.NewArgonHasher().CompareHashAndPassword(hash, psw)
}
