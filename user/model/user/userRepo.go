package user

import "context"

type Repository interface {
	Create(ctx context.Context, user User) (string, error)
	GetUUID(ctx context.Context, login, passwordHash string) (string, error)
	//TODO implement GetUser repo method
	//GetUser(ctx context.Context, uuid string) (GettingUserDTO, customErors)
}
