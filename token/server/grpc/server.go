package grpc

import (
	"fmt"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"github.com/viciousvs/OAuth-services/token/config"
	"github.com/viciousvs/OAuth-services/token/model/token/repository"
	"github.com/viciousvs/OAuth-services/token/service/tokenJWT"
	"google.golang.org/grpc"
	"log"
	"net"
)

//Server
type Server struct {
	repo repository.Repository
	jwt  tokenJWT.JWT
	tokenPb.UnimplementedTokenServiceServer
}

//NewServer
func NewServer(repo repository.Repository, jwt tokenJWT.JWT) *Server {
	return &Server{repo: repo, jwt: jwt}
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
	tokenPb.RegisterTokenServiceServer(srv, s)
	if err := srv.Serve(listener); err != nil {
		return fmt.Errorf("cannot serve GRPC server: %w", err)
	}

	return nil
}
