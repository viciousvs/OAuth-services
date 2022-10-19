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
func (h handler) Handle(ctx context.Context, req *userPb.LoginRequest) (user.User, error) {
	if req == nil {
		return user.User{}, customErors.ErrNilRequest
	}

	u, err := h.repo.GetOnlyByLogin(ctx, req.GetLogin())
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}
