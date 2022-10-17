package grpc

import (
	"fmt"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"github.com/viciousvs/OAuth-services/token/config"
	"google.golang.org/grpc"
	"log"
	"net"
)

//Server
type Server struct {
	tokenPb.UnimplementedTokenServiceServer
}

//NewServer
func NewServer() *Server {
	return &Server{}
}

//Run
func (s *Server) Run(cfg config.ServerConfig) error {
	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return fmt.Errorf("cannot listen %v: %w", cfg.Addr, err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Printf("listener close with customErors: %v", err)
		}
	}()

	srv := grpc.NewServer()
	tokenPb.RegisterTokenServiceServer(srv, s)
	if err := srv.Serve(listener); err != nil {
		return fmt.Errorf("cannot serve GRPC server: %w", err)
	}

	return nil
}
