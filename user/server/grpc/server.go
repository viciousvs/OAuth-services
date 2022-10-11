package grpc

import (
	"fmt"
	"github.com/viciousvs/OAuth-services/proto/userService"
	"github.com/viciousvs/OAuth-services/user/config"
	"github.com/viciousvs/OAuth-services/user/model/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	repo user.Repository
	userPb.UnimplementedUserServiceServer
}

func NewServer(repo user.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) Run(cfg config.ServerConfig) error {
	listener, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return fmt.Errorf("cannot listen %v: %w", cfg.Addr, err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Printf("listener close with error: %v", err)
		}
	}()

	srv := grpc.NewServer()
	userPb.RegisterUserServiceServer(srv, s)
	if err := srv.Serve(listener); err != nil {
		return fmt.Errorf("cannot serve GRPC server: %w", err)
	}

	return nil
}

func (s *Server) ShutDown() {
	log.Println("GRPC server shut down")
	s.ShutDown()
}
