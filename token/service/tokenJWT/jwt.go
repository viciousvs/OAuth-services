package tokenJWT

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/viciousvs/OAuth-services/token/model/token"
	"github.com/viciousvs/OAuth-services/token/utils/customErrors"
	"github.com/viciousvs/OAuth-services/user/utils/uuid"
	"io/ioutil"
	"strings"
	"time"
)

const (
	ACCESSTTL  = 30 * time.Minute    //30minute
	REFRESHTTL = 24 * time.Hour * 30 //30days
	ACCPREFIX  = "ACC"
	REFPREFIX  = "REF"
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
	if !uuid.IsValidUUID(userUUID) {
		return nil, customErrors.ErrInvalidUUID
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(j.prvKey)
	if err != nil {
		return nil, err
	}

	now, aID, rID := time.Now().UTC(), uuid.GenerateUUID(), uuid.GenerateUUID()

	aClaims := token.AccessTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(j.attl).UTC()),
			ID:        ACCPREFIX + aID,
			Subject:   userUUID,
		},
	}
	aTokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, aClaims).SignedString(key)
	if err != nil {
		return nil, err
	}

	accessToken := token.AccessTokenModel{TokenString: aTokenString, Claims: aClaims}

	rClaims := token.RefreshTokenClaims{
		AccessTokenClaims: &aClaims,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        REFPREFIX + rID,
			ExpiresAt: jwt.NewNumericDate(now.Add(j.rttl).UTC()),
			Subject:   userUUID,
		},
	}
	rTokenString, err := jwt.NewWithClaims(jwt.SigningMethodRS256, rClaims).SignedString(key)
	if err != nil {
		return nil, err
	}

	refreshToken := token.RefreshTokenModel{
		TokenString: rTokenString,
		Claims:      rClaims,
	}
	return &token.Tokens{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (j JWT) ValidateAccessToken(acctoken string) (string, string, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.pubKey)
	if err != nil {
		return "", "", customErrors.ErrorWithCode500(err.Error())
	}

	tok, err := jwt.Parse(acctoken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpecred method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return "", "", customErrors.ErrorWithCode400(err.Error())
	}

	accClaims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", "", customErrors.ErrInvalidAccessToken
	}
	js, err := json.Marshal(accClaims)
	if err != nil {
		return "", "", customErrors.ErrInvalidAccessToken
	}
	var aClaims token.AccessTokenClaims
	if err := json.Unmarshal(js, &aClaims); err != nil {
		return "", "", customErrors.ErrInvalidAccessToken
	}
	if !strings.HasPrefix(aClaims.ID, ACCPREFIX) {
		return "", "", customErrors.ErrInvalidAccessToken
	}
	return aClaims.Subject, aClaims.ID, nil
}

func (j JWT) Refresh(refToken string) (string, string, string, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(j.pubKey)
	if err != nil {
		return "", "", "", customErrors.ErrorWithCode500(err.Error())
	}

	tok, err := jwt.Parse(refToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpecred method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return "", "", "", customErrors.ErrorWithCode400(err.Error())
	}

	refreshClaims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", "", "", customErrors.ErrInvalidRefreshToken
	}
	js, err := json.Marshal(refreshClaims)
	if err != nil {
		return "", "", "", customErrors.ErrInvalidRefreshToken
	}
	var rClaims token.RefreshTokenClaims
	if err := json.Unmarshal(js, &rClaims); err != nil {
		return "", "", "", customErrors.ErrInvalidRefreshToken
	}
	if !strings.HasPrefix(rClaims.ID, REFPREFIX) {
		return "", "", "", customErrors.ErrInvalidRefreshToken
	}
	return rClaims.AccessTokenClaims.Subject, rClaims.AccessTokenClaims.ID, rClaims.ID, nil
}
