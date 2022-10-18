package service

import "github.com/viciousvs/OAuth-services/token/model/token"

type Service interface {
	Generate(userUUID string) (*token.Tokens, error)
	ValidateAccessToken(acctoken string) (string, error)
	Refresh(refToken string) (userUUID, aID, rID string, err error)
}
