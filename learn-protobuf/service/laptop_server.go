package service

import (
	"context"
	pb "learn-protobuf/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer用于生成laptop
type LaptopServer struct {
}

func NewLaptopServer() *LaptopServer {

	return &LaptopServer{}
}

// 生成laptop的service需要实现LaptopServiceServer接口
func (server *LaptopServer) CreateLaptop(
	cxt context.Context,
	req *pb.CreateLaptopRequest) (
	*pb.CreateLaptopResponse, error) {
	//生成一个laptop
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// validate uuid, uuid是client提供的，让service去生成相应的laptop, 也可以不提供，此时服务端来生成uuid
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not valid UUID : %v", laptop.Id)
		}
	} else {
		// clinet未提供uuid，服务端生成
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "can not generate a new laptop ID: %v", err)

		}
		laptop.Id = id.String()
	}
}
