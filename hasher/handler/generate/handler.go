package generate

import (
	"github.com/viciousvs/OAuth-services/hasher/service/argon2"
	"github.com/viciousvs/OAuth-services/hasher/utils/customErors"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
)

func Handle(req *hasherPb.Password) (string, error) {
	if req == nil {
		return "", customErors.ErrNilRequest
	}
	password := req.GetValue()
	if len(password) < 8 {
		return "", customErors.ErrInvalidData
	}

	return argon2.NewArgonHasher().GenerateHash(password)
}
