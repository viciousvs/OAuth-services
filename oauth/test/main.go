package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand"
	"os"
	"sync"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

const accToken = ""

func init() {
	if err := godotenv.Load("./../oauth.env"); err != nil {
		log.Printf("cannot load oauth.env file=> %v", err)
		log.Fatalf("used default values for config")
	}
}
func main() {

	successAccessResponses := 0
	successSUPresponses := 0
	successSINresponses := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			conn, err := grpc.Dial(os.Getenv("SERVER_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Println(err)
				return
			}
			defer func() {
				conn.Close()
				wg.Done()
			}()

			client := oauthPb.NewOAuthServiceClient(conn)
			_, err = client.SignUp(
				context.Background(),
				&oauthPb.SignUpRequest{
					Login:    randSeq(6),
					Password: randSeq(10),
				})
			if err != nil {
				log.Println(err)
				return
			}
			mu.Lock()
			successSUPresponses++
			mu.Unlock()
		}()
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			conn, err := grpc.Dial(os.Getenv("SERVER_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return
			}
			defer func() {
				conn.Close()
				wg.Done()
			}()

			client := oauthPb.NewOAuthServiceClient(conn)
			_, err = client.SignIn(
				context.Background(),
				&oauthPb.SignInRequest{
					Login:    "admin123",
					Password: "admin12345",
				})
			if err != nil {
				return
			}
			mu.Lock()
			successSINresponses++
			mu.Unlock()
		}()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(num int) {
			conn, err := grpc.Dial(os.Getenv("SERVER_ADDRESS"), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				//log.Printf("cannot connect to server iter:%d", num)
				return
			}
			defer func() {
				conn.Close()
				wg.Done()
			}()

			cl := oauthPb.NewOAuthServiceClient(conn)
			_, err = cl.Access(context.Background(), &oauthPb.AccessRequest{AccessToken: accToken})
			if err != nil {
				//log.Printf("cannot do access request to server iter:%d", num)
				return
			}
			mu.Lock()
			successAccessResponses++
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(successAccessResponses)

	fmt.Println(successSUPresponses)
	fmt.Println(successSINresponses)
}
