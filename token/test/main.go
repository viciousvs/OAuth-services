package main

import (
	"context"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot get a conn", err)
	}
	cl := tokenPb.NewTokenServiceClient(conn)
	resp, err := cl.GenerateTokens(context.Background(), &tokenPb.GenerateTokensRequest{Uuid: "af7c1fe6-d669-414e-b066-e9733f0de7a8"})
	if err != nil {
		log.Fatal("cannot do a generate tokens resp", err)
	}
	log.Println(resp)
}
