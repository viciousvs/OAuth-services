package user

import "context"

type Repository interface {
	Create(ctx context.Context, user User) (string, error)
	Get(ctx context.Context, login, passwordHash string) (string, error)
}
