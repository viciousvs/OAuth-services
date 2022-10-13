package grpc

import (
	"fmt"
	userPb "github.com/viciousvs/OAuth-services/proto/userService"
	"github.com/viciousvs/OAuth-services/user/config"
	"github.com/viciousvs/OAuth-services/user/model/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

//Server
type Server struct {
	repo user.Repository
	userPb.UnimplementedUserServiceServer
}

//NewServer
func NewServer(repo user.Repository) *Server {
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
	userPb.RegisterUserServiceServer(srv, s)
	if err := srv.Serve(listener); err != nil {
		return fmt.Errorf("cannot serve GRPC server: %w", err)
	}

	return nil
}
