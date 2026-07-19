package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	hellov1 "github.com/homepunks/zymyran/gen/go/hello/v1"
)

func main() {
	name := flag.String("name", "World", "name to greet")
	flag.Parse()

	conn, err := grpc.NewClient("127.0.0.1:6969", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	defer conn.Close()

	client := hellov1.NewGreeterServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Greet(ctx, &hellov1.GreetRequest{Name: *name})
	if err != nil {
		log.Fatalf("greet failed: %v", err)
	}

	log.Printf("response: %s", resp.GetMessage())
}
