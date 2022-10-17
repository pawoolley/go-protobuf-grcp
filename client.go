package main

import (
	"context"
	pb "go-echo/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:65535", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloWorldClient(conn)

	stream, err := client.SayHi(context.Background(), &pb.Request{Name: "Paul"})
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}

	for {
		response, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Stream closed")
				return
			} else {
				log.Fatalf("stream.Recv() failed: %v", err)
			}
		}
		log.Println(response.Message)
	}
}
