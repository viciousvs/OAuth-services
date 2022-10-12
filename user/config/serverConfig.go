package config

const srvAddr = "SERVER_ADDRESS"

//ServerConfig
type ServerConfig struct {
	Addr string
}

//MakeServerConfig
func MakeServerConfig() ServerConfig {
	return ServerConfig{
		Addr: getEnv(srvAddr, "localhost:50051"),
	}
}
