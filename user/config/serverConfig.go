package config

const srvAddr = "SERVER_ADDRESS"

type ServerConfig struct {
	Addr string
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		Addr: getEnv(srvAddr, "localhost:50051"),
	}
}
