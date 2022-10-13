package generate

import (
	"github.com/viciousvs/OAuth-services/hasher/model/hasher"
	"github.com/viciousvs/OAuth-services/hasher/utils/customErors"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
)

type handler struct {
	repo hasher.Repository
}

func NewHandler(repo hasher.Repository) *handler {
	return &handler{repo: repo}
}

func (h handler) Handle(req *hasherPb.Password) (string, error) {
	if req == nil {
		return "", customErors.ErrNilRequest
	}
	password := req.GetValue()
	if len(password) < 8 {
		return "", customErors.ErrInvalidData
	}

	return h.repo.GenerateHash(password)
}
