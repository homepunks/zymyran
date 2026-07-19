package main

import (
	"fmt"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	hellov1 "github.com/homepunks/zymyran/gen/go/hello/v1"
)

type greeterServer struct {
	hellov1.UnimplementedGreeterServiceServer
}

func (gs *greeterServer) Greet(ctx context.Context, req *hellov1.GreetRequest) (*hellov1.GreetResponse, error) {
	msg := fmt.Sprintf("Hello, %s!", req.GetName())

	return &hellov1.GreetResponse{Message: msg}, nil
}


func main() {
	conn, err := net.Listen("tcp", "127.0.0.1:6969")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	hellov1.RegisterGreeterServiceServer(s, &greeterServer{})

	log.Printf("listening on %s", conn.Addr())
	if err := s.Serve(conn); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}	
}
