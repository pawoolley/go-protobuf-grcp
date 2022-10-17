package main

import (
	"fmt"
	"go-echo/generated"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type helloWorldServer struct {
	pb.UnimplementedHelloWorldServer
}

func (s *helloWorldServer) SayHi(req *pb.Request, server pb.HelloWorld_SayHiServer) error {
	log.Printf("Received: %v\n", req.Name)

	for i := 0; i < 5; i++ {
		server.Send(&pb.Response{
			Message: fmt.Sprintf("Hi %v", req.Name),
		})
		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	lis, _ := net.Listen("tcp", "localhost:65535")
	grpcServer := grpc.NewServer()
	pb.RegisterHelloWorldServer(grpcServer, &helloWorldServer{})
	log.Println("Listening...")
	grpcServer.Serve(lis)
}
