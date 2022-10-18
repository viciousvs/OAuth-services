package tokenJWT

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/viciousvs/OAuth-services/token/model/token"
	"github.com/viciousvs/OAuth-services/user/utils/uuid"
	"io/ioutil"
	"log"
	"time"
)

const (
	ACCESSTTL  = 30 * time.Minute    //30minute
	REFRESHTTL = 24 * time.Hour * 30 //30days
)

type JWT struct {
	prvKey,
	pubKey []byte
	attl time.Duration
	rttl time.Duration
}

func NewJWT() (JWT, error) {
	prvKey, err := ioutil.ReadFile("cert/private.key")
	if err != nil {
		return JWT{}, err
	}
	pubKey, err := ioutil.ReadFile("cert/public.key")
	if err != nil {
		return JWT{}, err
	}

	return JWT{prvKey: prvKey, pubKey: pubKey, attl: ACCESSTTL, rttl: REFRESHTTL}, nil
}

func (j JWT) Generate(userUUID string) (*token.Tokens, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.prvKey)
	if err != nil {
		return nil, err
	}

	now, aID, rID := time.Now().UTC(), uuid.GenerateUUID(), uuid.GenerateUUID()

	aClaims := token.AccessTokenClaims{
		UserUUID: userUUID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(j.attl).UTC()),
			ID:        aID,
		},
	}
	aTokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, aClaims).SignedString(key)
	if err != nil {
		return nil, err
	}

	accessToken := token.AccessTokenModel{Id: aID, TokenString: aTokenString, Claims: aClaims}

	rClaims := token.RefreshTokenClaims{
		AccessTokenClaims: aClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        rID,
			ExpiresAt: jwt.NewNumericDate(now.Add(j.rttl).UTC()),
		},
	}
	rTokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, rClaims).SignedString(key)
	if err != nil {
		return nil, err
	}

	refreshToken := token.RefreshTokenModel{
		Id:            rID,
		AccessTokenId: aID,
		TokenString:   rTokenString,
		Claims:        rClaims,
	}
	return &token.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (j JWT) ValidateAccessToken(acctoken string) (string, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.pubKey)
	if err != nil {
		return "", err
	}

	tok, err := jwt.Parse(acctoken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpecred method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return "", err
	}
	log.Println(tok)
	log.Println(tok.Claims)
	claims, ok := tok.Claims.(token.AccessTokenClaims)
	if !ok || !tok.Valid {
		return "", fmt.Errorf("access token not valid")
	}
	return claims.UserUUID, nil
}

func (j JWT) Refresh(refToken string) (uuid, aID string, rID string, err error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.pubKey)
	if err != nil {
		return "", "", "", err
	}

	tok, err := jwt.Parse(refToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpecred method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return "", "", "", err
	}
	refreshClaims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", "", "", fmt.Errorf("refresh token not valid")
	}
	log.Println(ok)
	log.Println(refreshClaims)

	//return refreshClaims.AccessTokenClaims.UserUUID, refreshClaims.ID, refreshClaims.AccessTokenClaims.ID, nil
	return "", "", "", nil
}
