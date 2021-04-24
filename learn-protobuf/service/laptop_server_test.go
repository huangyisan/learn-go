package service_test

import (
	"context"
	pb "learn-protobuf/pb"
	"learn-protobuf/sample"
	"learn-protobuf/service"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()
	// 不提供uuid
	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""
	// 提供错误uuid
	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"
	// 重复uuid, 此时先将laptopDuplicateID存进去,让store里面已经存在laptopDuplicateID的数据
	laptopDuplicateID := sample.NewLaptop()
	storeDulplicateID := service.NewInMemoryLaptopStore()
	err := storeDulplicateID.Save(laptopDuplicateID)
	require.Nil(t, err)

	// 使用表驱动测试多种不同情况
	testCases := []struct {
		name   string
		laptop *pb.Laptop
		// 存储
		store service.LaptopStore
		// 预期状态码
		code codes.Code
	}{
		{
			// uuid由客户端生成
			name:   "success_with_id",
			laptop: sample.NewLaptop(),
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			// 不提供uuid
			name:   "success_no_id",
			laptop: laptopNoID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.OK,
		},
		{
			// 提供错误uuid
			name:   "failure_invalid_id",
			laptop: laptopInvalidID,
			store:  service.NewInMemoryLaptopStore(),
			code:   codes.InvalidArgument,
		},
		{
			// 重复uuid报错
			name:   "failure_duplicate_uuid",
			laptop: laptopDuplicateID,
			store:  storeDulplicateID,
			code:   codes.AlreadyExists,
		},
	}
	// 进行测试
	for i := range testCases {
		// 将测试用例保存到本地变量,避免并发问题
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// 创建一个 创建laptop的请求
			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}
			// 创建server
			server := service.NewLaptopServer(tc.store)
			// 发送请求
			res, err := server.CreateLaptop(context.Background(), req)
			// 返回的code为ok情况
			if tc.code == codes.OK {
				// 确保没有error返回,为nil
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEmpty(t, res.Id)
				// 如果发送的时候id大于0,也就是客户端提供uuid的时候,需要验证是否和服务端返回的id一致
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				// 返回code不为ok的情况
				// 确保有err返回
				require.Error(t, err)
				// 确保res返回为nil
				require.Nil(t, res)
				// 确保error是由this package or has a method `GRPCStatus() *Status`. 而非其他异常导致的.
				st, ok := status.FromError(err)
				require.True(t, ok)
				//期望的code和结果的code一致
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
