package compare

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

func (h handler) Handle(req *hasherPb.CompareRequest) (string, error) {
	if req == nil {
		return "", customErors.ErrNilRequest
	}
	hash, psw := req.GetPasswordHash(), req.GetPassword()
	if ok := validate(hash, psw); !ok {
		return "", customErors.ErrInvalidData
	}

	if err := h.repo.CompareHashAndPassword(hash, psw); err != nil {
		return "", err
	}

	return hash, nil
}
