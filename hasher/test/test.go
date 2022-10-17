package main

import (
	"context"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	wg := sync.WaitGroup{}
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := hasherPb.NewHasherServiceClient(conn)
			hash, err := client.GenerateHash(context.Background(), &hasherPb.Password{Value: "NewPassword"})
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(hash)
		}()
	}
	wg.Wait()

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := hasherPb.NewHasherServiceClient(conn)
			_, err := client.CompareHashAndPassword(context.Background(), &hasherPb.CompareRequest{
				PasswordHash: "$argon2id$v=19$m=65536,t=1,p=2$1TAXRA0+arO2+E7y6oG1IQ$EHNMnsOkEdMmVsuhAi2gbZOCXXxq4Cmpji8MbhK7C0g",
				Password:     "wafawfasasc",
			})
			if err != nil {
				log.Println(err)
				return
			}
		}()
	}

	wg.Wait()

}
