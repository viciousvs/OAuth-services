package getUUID

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
func (h handler) Handle(ctx context.Context, req *userPb.GetUUIDRequest) (string, error) {
	// TODO validation
	if req == nil {
		//TODO come up customErors
		return "", customErors.ErrNilRequest
	}
	// TODO logic check

	//has post with uuid from request
	//h.repo.HasById // bool
	u := makeGetUUIDDTO(req.GetLogin(), req.GetPassword())

	uuid, err := h.repo.GetUUID(ctx, u.Login, u.PasswordHash)
	if err != nil {
		//TODO create utils ERRORS, come up
		return "", err
	}
	return uuid, nil
}
