package user

import "context"

type Repository interface {
	Create(ctx context.Context, user User) (string, error)
	GetUser(ctx context.Context, login string) (User, error)
	//TODO implement GetUser repo method
	//GetUser(ctx context.Context, uuid string) (GettingUserDTO, customErors)
}
