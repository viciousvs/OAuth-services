package grpc

import (
	"fmt"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"github.com/viciousvs/OAuth-services/token/config"
	"github.com/viciousvs/OAuth-services/token/model/token/repository"
	"github.com/viciousvs/OAuth-services/token/model/token/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

//Server
type Server struct {
	repo    repository.Repository
	service service.Service
	tokenPb.UnimplementedTokenServiceServer
}

//NewServer
func NewServer(repo repository.Repository, service service.Service) *Server {
	return &Server{repo: repo, service: service}
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
