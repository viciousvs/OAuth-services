package grpc

import (
	"fmt"
	"github.com/viciousvs/OAuth-services/hasher/config"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
	"google.golang.org/grpc"
	"log"
	"net"
)

//Server
type Server struct {
	hasherPb.UnimplementedHasherServiceServer
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
			log.Printf("listener close with customErrors: %v", err)
		}
	}()

	srv := grpc.NewServer()
	hasherPb.RegisterHasherServiceServer(srv, s)
	if err := srv.Serve(listener); err != nil {
		return fmt.Errorf("cannot serve GRPC server: %w", err)
	}

	return nil
}
