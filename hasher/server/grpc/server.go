package grpc

import (
	"fmt"
	"github.com/viciousvs/OAuth-services/hasher/config"
	"github.com/viciousvs/OAuth-services/hasher/model/hasher"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
	"google.golang.org/grpc"
	"log"
	"net"
)

//Server
type Server struct {
	repo hasher.Repository
	hasherPb.UnimplementedHasherServiceServer
}

//NewServer
func NewServer(repo hasher.Repository) *Server {
	return &Server{repo: repo}
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
	hasherPb.RegisterHasherServiceServer(srv, s)
	if err := srv.Serve(listener); err != nil {
		return fmt.Errorf("cannot serve GRPC server: %w", err)
	}

	return nil
}
