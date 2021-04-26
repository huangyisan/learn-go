package service_test

import (
	"context"
	pb "learn-protobuf/pb"
	"learn-protobuf/sample"
	"learn-protobuf/service"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestClientCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopServer, serverAddress := startTestLaptopServer(t)
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
	other, err := laptopServer.Store.Find(res.Id)
	require.NoError(t, err)
	require.NotNil(t, other)

	// check the saved laptop is the same as the one we send
	requireSameLaptop(t, laptop, other)

}

// start grpc server, 返回laptopServer和监听地址信息
func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	// create grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	// 0表示随机分配端口
	listener, err := net.Listen("tcp", "127.0.0.1:8880")
	require.NoError(t, err)
	// 将他运行在go协程里,以至于下面的代码不会被block
	go grpcServer.Serve(listener)
	return laptopServer, listener.Addr().String()
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
