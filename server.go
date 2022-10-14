package main

import (
	"context"
	"fmt"
	"go-echo/generated"
	"google.golang.org/grpc"
	"log"
	"net"
)

type helloWorldServer struct {
	pb.UnimplementedHelloWorldServer
}

func (s *helloWorldServer) SayHi(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v\n", req.Name)
	return &pb.Response{
		Message: fmt.Sprintf("Hi, %v", req.Name),
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", "localhost:65535")
	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServer, &helloWorldServer{})
	log.Println("Listening...")
	grpcServer.Serve(lis)
}
