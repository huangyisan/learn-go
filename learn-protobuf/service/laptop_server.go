package service

import (
	"bytes"
	"context"
	"errors"
	"io"
	pb "learn-protobuf/pb"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 最大图片大小 1MB
const maxImageSize = 1 << 20

// LaptopServer用于生成laptop
type LaptopServer struct {
	laptopStore LaptopStore
	imageStore  ImageStore

	pb.UnimplementedLaptopServiceServer
}

func NewLaptopServer(laptopStore LaptopStore, imageStore ImageStore) *LaptopServer {

	return &LaptopServer{laptopStore, imageStore, pb.UnimplementedLaptopServiceServer{}}
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

	// time.Sleep(time.Second * 6)

	// check context err on server
	if err := contextError(ctx); err != nil {
		return nil, err
	}

	// save to db
	err := server.laptopStore.Save(laptop)
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

func (server *LaptopServer) SearchLaptop(req *pb.SearchLapTopRequest, stream pb.LaptopService_SearchLaptopServer) error {
	filter := req.GetFilter()
	log.Printf("receive a search-laptop request with filter: %v", filter)
	err := server.laptopStore.Search(
		// get context from stream, and pass to func Search
		stream.Context(),
		filter,
		// 这里的laptop被传入的是符合filter条件的laptop
		func(laptop *pb.Laptop) error {
			res := &pb.SearchLapTopResponse{
				Laptop: laptop,
			}
			err := stream.Send(res)
			if err != nil {
				return err
			}

			log.Printf("sent laptop with id: %s", laptop.Id)
			return nil
		},
	)
	if err != nil {
		return status.Errorf(codes.Internal, "unexcepted error: %v", err)
	}
	return nil
}

// uploadimage 客户端发送信息给服务端，服务端如何进行处理
func (server *LaptopServer) UploadImage(stream pb.LaptopService_UploadImageServer) error {
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "can not receive image info"))
	}

	laptopID := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("receive an upload-image request for laptop %s with image type %s", laptopID, imageType)

	// 确保laptopID存在
	laptop, err := server.laptopStore.Find(laptopID)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
	}
	if laptop == nil {
		return logError(status.Errorf(codes.InvalidArgument, "laptop %s doesn't exist", laptopID))
	}

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		// 每一次循环之前，都检查context是否有error
		if err := contextError(stream.Context()); err != nil {
			return err
		}
		log.Print("waiting to receive more data")
		req, err := stream.Recv()
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)

		imageSize += size

		if imageSize > maxImageSize {
			return logError(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize))
		}

		_, err = imageData.Write(chunk)

		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	imageID, err := server.imageStore.Save(laptopID, imageType, imageData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save image to the store :%v", err))
	}

	// create a uploadImageResponse
	res := &pb.UploadImageResponse{
		Id:   imageID,
		Size: uint32(imageSize),
	}

	// send response
	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "can not send upload image response :%v", err))
	}
	// save successful
	log.Printf("saved image with id: %s, size: %d", imageID, imageSize)
	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "canceled by client"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
