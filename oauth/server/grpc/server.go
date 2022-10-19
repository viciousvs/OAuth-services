package grpc

import (
	"fmt"
	"github.com/viciousvs/OAuth-services/oauth/config"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
	"google.golang.org/grpc"
	"log"
	"net"
)

//Server
type Server struct {
	oauthPb.UnimplementedOAuthServiceServer
}

//NewServer
func NewServer() *Server {
	return &Server{}
}

//Run
func (s *Server) Run(cfg config.ServerConfig) error {
	if s == nil {
		return fmt.Errorf("nil server or services")
	}
	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return fmt.Errorf("cannot listen %v: %w", cfg.Addr, err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Printf("listener close with customErrors: %v", err)
		}
	}()

	srv := grpc.NewServer()
	oauthPb.RegisterOAuthServiceServer(srv, s)
	if err := srv.Serve(listener); err != nil {
		return fmt.Errorf("cannot serve GRPC server: %w", err)
	}

	return nil
}
