package getUUID

import (
	"context"
	userPb "github.com/viciousvs/OAuth-services/proto/userService"
	"github.com/viciousvs/OAuth-services/user/model/user"
	"github.com/viciousvs/OAuth-services/user/utils/customErors"
	"github.com/viciousvs/OAuth-services/user/utils/hashPassword"
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
func (h handler) Handle(ctx context.Context, req *userPb.GetUUIDRequest) (string, error) {
	if req == nil {
		//TODO come up customErors
		return "", customErors.ErrNilRequest
	}

	u := makeGetUserDTO(req.GetLogin(), req.GetPassword())
	if err := u.Validate(); err != nil {
		return "", customErors.ErrInvalidData
	}

	userRepo, err := h.repo.GetUser(ctx, u.Login)
	if err != nil {
		//TODO create utils ERRORS, come up
		return "", err
	}

	if empty(userRepo.UUID) || empty(userRepo.Login) || empty(userRepo.PasswordHash) {
		return "", customErors.ErrEmptyBody
	}

	err = hashPassword.CompareHashAndPassword(userRepo.PasswordHash, u.Password)
	if err != nil {
		return "", customErors.ErrIncorrectPasswordOrLogin
	}

	return userRepo.UUID, nil
}
