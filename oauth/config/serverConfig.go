package config

const srvAddr = "SERVER_ADDRESS"

//ServerConfig efsef
type ServerConfig struct {
	Addr string
}

//MakeServerConfig fef
func MakeServerConfig() ServerConfig {
	return ServerConfig{
		Addr: GetEnv(srvAddr, "localhost:50054"),
	}
}
