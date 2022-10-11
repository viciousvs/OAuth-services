package user

import (
	"context"
	"github.com/viciousvs/OAuth-services/user/storage/postgres"
)

type repoPostgres struct {
	db *postgres.PostgresDB
}

func NewPostgresRepo(db *postgres.PostgresDB) Repository {
	return repoPostgres{db: db}
}
func (r repoPostgres) Create(ctx context.Context, user User) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r repoPostgres) Get(ctx context.Context, login, passwordHash string) (string, error) {
	//TODO implement me
	panic("implement me")
}
