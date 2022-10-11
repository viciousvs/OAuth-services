package config

const (
	pgUser = "POSTGRES_USER"
	pgPass = "POSTGRES_PASSWORD"
	pgHost = "POSTGRES_HOST"
	pgPort = "POSTGRES_PORT"
	pgDB   = "POSTGRES_DB"
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
		User:         getEnv(pgUser, "postgres"),
		Password:     getEnv(pgPass, "postgres"),
		Host:         getEnv(pgHost, "localhost"),
		Port:         getEnv(pgPort, "5432"),
		DatabaseName: getEnv(pgDB, "user"),
	}
}
