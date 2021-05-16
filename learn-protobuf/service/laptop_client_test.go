package service_test

import (
	"bufio"
	"context"
	"fmt"
	"io"
	pb "learn-protobuf/pb"
	"learn-protobuf/sample"
	"learn-protobuf/service"
	"net"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopStore := service.NewInMemoryLaptopStore()
	serverAddress := startTestLaptopServer(t, laptopStore, nil)
	laptopClient := newTestLaptopClient(t, serverAddress)

	// 创建laptop
	laptop := sample.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	res, err := laptopClient.CreateLaptop(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, expectedID, res.Id)

	// check laptop saved in store
	other, err := laptopStore.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check the saved laptop is the same as the one we send
	requireSameLaptop(t, laptop, other)

}

func TestClientSearchLaptop(t *testing.T) {
	t.Parallel()
	filter := &pb.Filter{
		MaxPriceUsd: 2000,
		MinCpuCores: 4,
		MinCpuGhz:   2.2,
		MinRam: &pb.Memory{
			Value: 8,
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	laptopStore := service.NewInMemoryLaptopStore()

	expectedIDs := make(map[string]bool)

	for i := 0; i < 6; i++ {
		laptop := sample.NewLaptop()

		switch i {
		case 0:
			laptop.PriceUsd = 2500
		case 1:
			laptop.Cpu.NumberCores = 2
		case 2:
			laptop.Cpu.MinGhz = 2.0
		case 3:
			laptop.Ram = &pb.Memory{Value: 4096, Unit: pb.Memory_MEGABYTE}
		case 4:
			laptop.PriceUsd = 1999
			laptop.Cpu.NumberCores = 4
			laptop.Cpu.MinGhz = 2.5
			laptop.Cpu.MaxGhz = 4.5
			laptop.Ram = &pb.Memory{Value: 16, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		case 5:
			laptop.PriceUsd = 2000
			laptop.Cpu.NumberCores = 6
			laptop.Cpu.MinGhz = 2.8
			laptop.Cpu.MaxGhz = 5.0
			laptop.Ram = &pb.Memory{Value: 64, Unit: pb.Memory_GIGABYTE}
			expectedIDs[laptop.Id] = true
		}
		err := laptopStore.Save(laptop)
		require.NoError(t, err)
	}
	serverAddress := startTestLaptopServer(t, laptopStore, nil)
	laptopClient := newTestLaptopClient(t, serverAddress)

	req := &pb.SearchLapTopRequest{
		Filter: filter,
	}
	stream, err := laptopClient.SearchLaptop(context.Background(), req)
	require.NoError(t, err)

	found := 0
	for {
		// 这边接受到的res，已经在server端进行了过滤。只剩下符合的条件的laptop数据
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}
		fmt.Println(res.GetLaptop().GetId())
		require.NoError(t, err)
		require.Contains(t, expectedIDs, res.GetLaptop().GetId())
		fmt.Println(found)
		found += 1
	}
	fmt.Printf("%v", expectedIDs)
	require.Equal(t, len(expectedIDs), found)
}

func TestClientUploadImage(t *testing.T) {
	t.Parallel()

	const testImageFolder = "../tmp"

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore(testImageFolder)

	laptop := sample.NewLaptop()
	err := laptopStore.Save(laptop)
	require.NoError(t, err)

	serverAddress := startTestLaptopServer(t, laptopStore, imageStore)
	laptopClient := newTestLaptopClient(t, serverAddress)

	imagePath := fmt.Sprintf("%s/laptop.jpg", testImageFolder)
	file, err := os.Open(imagePath)
	require.NoError(t, err)
	defer file.Close()

	stream, err := laptopClient.UploadImage(context.Background())
	require.NoError(t, err)
	imageType := filepath.Ext(imagePath)

	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laptop.GetId(),
				ImageType: imageType,
			},
		},
	}
	// 将请求发送给server
	err = stream.Send(req)
	require.NoError(t, err)
	// 以大块读取文件内容
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)
	size := 0

	for {
		// 返回字节数和err
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		require.NoError(t, err)
		size += n

		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}
		// 多次发送image data
		err = stream.Send(req)
		require.NoError(t, err)

	}
	// 关闭并接受请求，此时接受了服务端会返回的信息
	res, err := stream.CloseAndRecv()
	require.NoError(t, err)
	require.NotZero(t, res.GetId())
	require.EqualValues(t, size, res.GetSize())

	// 检查是否生成图片
	savedImagePath := fmt.Sprintf("%s/%s%s", testImageFolder, res.GetId(), imageType)
	require.FileExists(t, savedImagePath)
	// 测试完成后确保删除文件无报错
	err = os.Remove(savedImagePath)
	require.NoError(t, err)
}

// start grpc server, 返回laptopServer和监听地址信息
func startTestLaptopServer(t *testing.T, laptopStore service.LaptopStore, imageStore service.ImageStore) string {
	laptopServer := service.NewLaptopServer(laptopStore, imageStore)
	// create grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	// 0表示随机分配端口
	listener, err := net.Listen("tcp", ":0")
	require.NoError(t, err)
	// 将他运行在go协程里,以至于下面的代码不会被block
	go grpcServer.Serve(listener)
	return listener.Addr().String()
}

// 创建client端
func newTestLaptopClient(t *testing.T, serverAddress string) pb.LaptopServiceClient {
	// 仅用于测试,使用不安全的连接 grpc.WithInsecure()
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	require.NoError(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	// require.Equal(t, laptop1, laptop2)
	json1, err := protojson.Marshal(laptop1)
	require.NoError(t, err)
	json2, err := protojson.Marshal(laptop2)
	require.NoError(t, err)

	require.Equal(t, json1, json2)

}
