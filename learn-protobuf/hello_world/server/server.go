package main

import (
	"context"
	"google.golang.org/grpc"
	pb "learn-protobuf/hello_world/helloWorld"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedPrintHelloServer
}

const (
	port = ":50051"
)

func (s *server) PrintHelloWorld(ctx context.Context, in *pb.ReceiveMessage) (*pb.ResponseMessage, error) {
	return &pb.ResponseMessage{
		ResponseString: "hello world",
	}, nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPrintHelloServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
