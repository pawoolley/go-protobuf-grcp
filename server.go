package main

import (
	"context"
	"fmt"
	"go-echo/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type helloWorldServer struct {
	pb.UnimplementedHelloWorldServer
}

func (s *helloWorldServer) SayHi(_ context.Context, req *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v\n", req.Name)
	return &pb.Response{
		Message: fmt.Sprintf("Hi, %v", req.Name),
	}, nil
}

func main() {
	lis, _ := net.Listen("tcp", "localhost:65535")
	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServer, &helloWorldServer{})
	// Register reflection service on gRPC server. Without doing this, IJ's built-in HTTP client
	// won't do GRPC requests and fails with a "Reflective gRPC call timed out" error.
	reflection.Register(grpcServer)
	log.Println("Listening...")
	//goland:noinspection GoUnhandledErrorResult
	grpcServer.Serve(lis)
}
