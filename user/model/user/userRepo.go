package user

import "context"

type Repository interface {
	Create(ctx context.Context, user User) (string, error)
	GetByLogin(ctx context.Context, login, passwordHash string) (User, error)
	GetOnlyByLogin(ctx context.Context, login string) (User, error)
}
