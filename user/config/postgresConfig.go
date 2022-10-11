package config

const (
	PG_USER = "POSTGRES_USER"
	PG_PASS = "POSTGRES_PASSWORD"
	PG_HOST = "POSTGRES_HOST"
	PG_PORT = "POSTGRES_PORT"
	PG_DB   = "POSTGRES_DB"
)

type PostgresConfig struct {
	User,
	Password,
	Host,
	Port,
	DatabaseName string
}

func MakePostgresConfig() PostgresConfig {
	return PostgresConfig{
		User:         getEnv(PG_USER, "postgres"),
		Password:     getEnv(PG_PASS, "postgres"),
		Host:         getEnv(PG_HOST, "localhost"),
		Port:         getEnv(PG_PORT, "5432"),
		DatabaseName: getEnv(PG_DB, "user"),
	}
}
