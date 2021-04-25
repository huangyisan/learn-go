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

}

// start grpc server, 返回laptopServer和监听地址信息
func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopServer := service.NewLaptopServer(service.NewInMemoryLaptopStore())
	// create grpc server
	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)
	// 0表示随机分配端口
	listener, err := net.Listen("tcp", "0")
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
