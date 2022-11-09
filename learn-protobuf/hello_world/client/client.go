package main

import (
	"context"
	"google.golang.org/grpc"
	pb "learn-protobuf/hello_world/helloWorld"
	"log"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "abc"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewPrintHelloClient(conn)

	msg := defaultName
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	r, err := c.PrintHelloWorld(ctx, &pb.ReceiveMessage{ReceiveString: msg})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("output: %s", r.ResponseString)
}
