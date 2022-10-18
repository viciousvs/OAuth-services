package token

import (
	"github.com/golang-jwt/jwt/v4"
)

type Tokens struct {
	RefreshToken RefreshTokenModel `json:"refresh_token,omitempty"`
	AccessToken  AccessTokenModel  `json:"access_token,omitempty"`
}
type RefreshTokenModel struct {
	Id            string             `json:"id,omitempty"`
	AccessTokenId string             `json:"access_token_id,omitempty"`
	TokenString   string             `json:"token_string,omitempty"`
	Claims        RefreshTokenClaims `json:"claims,omitempty"`
}
type AccessTokenModel struct {
	Id          string            `json:"id,omitempty"`
	TokenString string            `json:"token_string,omitempty"`
	Claims      AccessTokenClaims `json:"claims,omitempty"`
}
type AccessTokenClaims struct {
	UserUUID             string `json:"user_uuid,omitempty"`
	jwt.RegisteredClaims `json:"jwt_registered_claims,omitempty"`
}
type RefreshTokenClaims struct {
	AccessTokenClaims    AccessTokenClaims `json:"access_token_claims"`
	jwt.RegisteredClaims `json:"jwt_registered_claims,omitempty"`
}
