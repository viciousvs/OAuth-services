package getUUID

//
//import (
//	"context"
//	"fmt"
//	userPb "github.com/viciousvs/OAuth-services/proto/userService"
//	"github.com/viciousvs/OAuth-services/user/model/user"
//)
//
//type handler struct {
//	repo user.Repository
//}
//
//func NewHandler(repo user.Repository) *handler {
//	return &handler{repo: repo}
//}
//
//func (h handler) Handle(ctx context.Context, req *userPb.GetUUIDRequest) (string, customErors) {
//	// TODO validation
//	if req == nil {
//		//TODO come up customErors
//		return "", fmt.Errorf("")
//	}
//	// TODO logic check
//
//	//has post with uuid from request
//	//h.repo.HasById // bool
//	u := MakeGetUUIDDTO(req.GetLogin(), req.GetPassword())
//
//	uuid, err := h.repo.GetUUID(ctx, u.Login, u.PasswordHash)
//	if err != nil {
//		//TODO create utils ERRORS, come up
//		return "", err
//	}
//	return uuid, nil
//}
