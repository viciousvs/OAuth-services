package create

import (
	"context"
	userPb "github.com/viciousvs/OAuth-services/proto/userService"
	"github.com/viciousvs/OAuth-services/user/model/user"
	"github.com/viciousvs/OAuth-services/user/utils/customErors"
)

//handler
type handler struct {
	repo user.Repository
}

//NewHandler
func NewHandler(repo user.Repository) *handler {
	return &handler{repo: repo}
}

//Handle
func (h handler) Handle(ctx context.Context, req *userPb.CreateRequest) (string, error) {
	// validation
	if req == nil {
		return "", customErors.ErrNilRequest
	}
	if len(req.GetPassword()) < 8 || len(req.GetLogin()) < 3 {
		return "", customErors.ErrInvalidData
	}

	u, err := makeCreateUserDTO(req.GetLogin(), req.GetPassword())
	if err != nil {
		return "", err
	}
	if err := u.Validate(); err != nil {
		return "", customErors.ErrInvalidData
	}

	createUser := user.MakeUser(u.UUID, u.Login, u.PasswordHash)
	uuid, err := h.repo.Create(ctx, createUser)
	if err != nil {
		return "", err
	}

	return uuid, nil
}
