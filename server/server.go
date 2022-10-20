package main

import (
	"fmt"
	"go-protobuf-grcp/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type helloWorldServer struct {
	protobuf.UnimplementedHelloWorldServer
}

func (s *helloWorldServer) SayHi(req *protobuf.Request, server protobuf.HelloWorld_SayHiServer) error {
	log.Printf("Received: %v\n", req.Name)

	for i := 0; i < 5; i++ {
		//goland:noinspection GoUnhandledErrorResult
		server.Send(&protobuf.Response{
			Message: fmt.Sprintf("Hi, %v", req.Name),
		})
		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	lis, _ := net.Listen("tcp", "localhost:65535")
	grpcServer := grpc.NewServer()
	protobuf.RegisterHelloWorldServer(grpcServer, &helloWorldServer{})
	// Register reflection service on gRPC server. Without doing this, IJ's built-in HTTP client
	// won't do GRPC requests and fails with a "Reflective gRPC call timed out" error.
	reflection.Register(grpcServer)
	log.Println("Listening...")
	//goland:noinspection GoUnhandledErrorResult
	grpcServer.Serve(lis)
}
