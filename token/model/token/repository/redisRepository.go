package repository

import (
	"context"
	"github.com/viciousvs/OAuth-services/token/model/token"
	"github.com/viciousvs/OAuth-services/token/service/tokenJWT"
	"github.com/viciousvs/OAuth-services/token/storage/redisRepo"
)

type redisRepository struct {
	db *redisRepo.RedisDB
}

func (r redisRepository) DeleteTokens(ctx context.Context, accessID, refreshID string) error {
	//TODO implement me
	return r.db.Del(ctx, accessID, refreshID).Err()
}

func (r redisRepository) SetTokens(ctx context.Context, tokens *token.Tokens) error {
	//TODO implement
	err := r.db.Set(ctx, tokens.RefreshToken.Id, tokens.RefreshToken.TokenString, tokenJWT.REFRESHTTL).Err()
	if err != nil {
		return err
	}

	err = r.db.Set(ctx, tokens.AccessToken.Id, tokens.AccessToken.TokenString, tokenJWT.ACCESSTTL).Err()
	if err != nil {
		return err
	}

	return nil
}

func NewRedisReporitory(db *redisRepo.RedisDB) redisRepository {
	return redisRepository{db: db}
}
