package repository

import (
	"context"
	"github.com/viciousvs/OAuth-services/token/model/token"
)

type Repository interface {
	SetTokens(ctx context.Context, tokens *token.Tokens) error
	DeleteTokens(ctx context.Context, accessID, refreshID string) error
}
