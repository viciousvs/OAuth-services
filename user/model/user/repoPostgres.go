package user

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/viciousvs/OAuth-services/user/storage/postgres"
	"github.com/viciousvs/OAuth-services/user/utils/customErors"
	"github.com/viciousvs/OAuth-services/user/utils/hashPassword"
)

type repoPostgres struct {
	db *postgres.PostgresDB
}

//NewPostgresRepo
func NewPostgresRepo(db *postgres.PostgresDB) Repository {
	return repoPostgres{db: db}
}

//Create
func (r repoPostgres) Create(ctx context.Context, user User) (string, error) {
	stmt := `INSERT INTO users(uuid, login, password_hash) VALUES ($1, $2, $3)`

	ct, err := r.db.Exec(ctx, stmt, user.UUID, user.Login, user.PasswordHash)
	if err != nil {
		if pgerr, ok := err.(*pgconn.PgError); ok {
			if pgerr.Code == customErors.PgCodeRecordAlreadyExists {
				return "", customErors.ErrRecordAlreadyExists
			}
			if pgerr.Code == customErors.PgCodeViolatesCheckConstraints {
				return "", customErors.ErrViolatesCheckConstraints
			}
		}
		return "", err
	}

	if ct.RowsAffected() != 1 {
		return "", customErors.ErrRowsAffected
	}
	return user.UUID, nil
}

//GetUUID
func (r repoPostgres) GetUUID(ctx context.Context, login, passwordHash string) (string, error) {
	stmt := `SELECT uuid, password_hash FROM users WHERE login=$1`

	var uuid, pHashFromDB string

	err := r.db.QueryRow(ctx, stmt, login).Scan(&uuid, &pHashFromDB)
	if err != nil {
		if err == pgx.ErrNoRows {
			return "", customErors.ErrNotFound
		}
		return "", err
	}

	//TODO compare hashes
	if !hashPassword.CompareHashes(passwordHash, pHashFromDB) {
		return "", customErors.ErrDifferentPasswordHashes
	}

	return uuid, nil
}

//func (r repoPostgres) GetUser(ctx context.Context, login, hashPassword string) (GettingUserDTO, error) {
//	//TODO implement me
//	panic("implement me")
//}
