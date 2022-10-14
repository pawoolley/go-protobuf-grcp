package main

import (
	"context"
	pb "go-echo/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:65535", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloWorldClient(conn)

	response, err := client.SayHi(context.Background(), &pb.Request{Name: "Paul"})
	if err != nil {
		log.Fatalf("client.GetFeature failed: %v", err)
	}
	log.Println(response.Message)
}
