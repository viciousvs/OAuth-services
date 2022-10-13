package getByLogin

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

//NewHandler constructor
func NewHandler(repo user.Repository) *handler {
	return &handler{repo: repo}
}

//Handle
func (h handler) Handle(ctx context.Context, req *userPb.GetByLoginRequest) (user.User, error) {
	if req == nil {
		return user.User{}, customErors.ErrNilRequest
	}
	dto := _DTO{
		Login:        req.GetLogin(),
		PasswordHash: req.GetPasswordHash(),
	}
	if err := dto.validate(); err != nil {
		return user.User{}, customErors.ErrInvalidData(err.Error())
	}

	u, err := h.repo.GetByLogin(ctx, dto.Login, dto.PasswordHash)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}
