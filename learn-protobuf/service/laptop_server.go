package service

import (
	"context"
	"errors"
	pb "learn-protobuf/pb"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// LaptopServer用于生成laptop
type LaptopServer struct {
	Store LaptopStore
	pb.UnimplementedLaptopServiceServer
}

func NewLaptopServer(store LaptopStore) *LaptopServer {

	return &LaptopServer{store, pb.UnimplementedLaptopServiceServer{}}
}

// 生成laptop的service需要实现LaptopServiceServer接口的CreateLaptop方法
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest) (
	*pb.CreateLaptopResponse, error) {
	//获取laptop,注意,这个laptop是客户端创建好发过来给server的
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// validate uuid, uuid是client提供的(默认会生成uuid)，让service去保存laptop, 也可以不提供，此时服务端来生成uuid
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

	time.Sleep(time.Second * 6)

	// check context err on server
	if ctx.Err() == context.Canceled {
		log.Print("canceled by client")
		return nil, status.Error(codes.Canceled, "canceled by client")
	}

	if ctx.Err() != context.DeadlineExceeded {
		log.Print("deadline is exceeded")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	}

	// save to db
	err := server.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		// 查看错误是否因为已经存在的id
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "can not save laptop to the memory store: %v", err)
	}
	log.Printf("Saved laptop with id: %s", laptop.Id)

	// 使用笔记本的id创建一个相应对象,返回给请求调用者
	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}
