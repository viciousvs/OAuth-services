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
	if req == nil {
		return "", customErors.ErrNilRequest
	}

	u := newDTOwithUUID(req.Login, req.PasswordHash)
	if err := u.validate(); err != nil {
		return "", customErors.ErrInvalidData(err.Error())
	}

	uuid, err := h.repo.Create(ctx, user.NewUser(u.UUID, u.Login, u.PasswordHash))
	if err != nil {
		return "", err
	}

	return uuid, nil
}
