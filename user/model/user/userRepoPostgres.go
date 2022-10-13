package user

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/viciousvs/OAuth-services/user/storage/postgres"
	"github.com/viciousvs/OAuth-services/user/utils/customErors"
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

//GetByLogin
func (r repoPostgres) GetByLogin(ctx context.Context, login, passwordHash string) (User, error) {
	stmt := `SELECT uuid, login, password_hash FROM users WHERE (login=$1 AND password_hash=$2)`

	var user User

	err := r.db.QueryRow(ctx, stmt, login, passwordHash).Scan(&user.UUID, &user.Login, &user.PasswordHash)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user, customErors.ErrNotFound
		}
		return user, err
	}

	return user, nil
}
