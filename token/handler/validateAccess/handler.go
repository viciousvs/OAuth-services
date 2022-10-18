package validateAccess

import (
	"context"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"github.com/viciousvs/OAuth-services/token/model/token/repository"
	"github.com/viciousvs/OAuth-services/token/model/token/service"
)

type handler struct {
	repo    repository.Repository
	service service.Service
}

func NewHandler(repo repository.Repository, service service.Service) *handler {
	return &handler{repo: repo, service: service}
}

func (h handler) Handle(ctx context.Context, request *tokenPb.ValidateAccessRequest) (string, error) {
	accToken := request.GetAccessToken()
	//TODO VALIDATE
	return h.service.ValidateAccessToken(accToken)
}
