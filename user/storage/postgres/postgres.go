package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/viciousvs/OAuth-services/user/config"
	"log"
	"sync"
)

var (
	once   sync.Once
	pgPool *PostgresDB
)

//PostgresDB
type PostgresDB struct {
	*pgxpool.Pool
}

// NewPostgresDB can end with log.Fatal()
func NewPostgresDB(cfg config.PostgresConfig) *PostgresDB {
	connUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName)
	once.Do(func() {
		var err error

		pgPool, err = getPostgresDB(connUrl)
		if err != nil {
			log.Fatal(err)
		}
	})
	return pgPool
}

// getPostgresDB
func getPostgresDB(connUrl string) (*PostgresDB, error) {
	ctx := context.TODO()

	pcfg, err := pgxpool.ParseConfig(connUrl)
	if err != nil {
		return nil, fmt.Errorf("cannot get connection pool: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pcfg)
	err = pool.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("cannot ping connection pool: %w", err)
	}

	return &PostgresDB{pool}, nil
}
