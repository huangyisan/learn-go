package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	pb "learn-protobuf/pb"
	"learn-protobuf/sample"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createLaptop(laptopClient pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// 超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("laptop already exists")
		} else {
			log.Fatal("can not create laptop ", err)
		}
		// 有报错直接return停止
		return
	}
	fmt.Printf("create new laptop with id: %s", res.Id)

}

func searchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter) {
	log.Print("search filter: ", filter)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.SearchLapTopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatal("can not search laptop: ", err)
	}
	// 使用for循环读取stream内容，直到读到EOF
	for {
		res, err := stream.Recv()
		// 读取完数据
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}
		laptop := res.GetLaptop()
		log.Print("- found: ", laptop.GetId())
		log.Print(" + brand: ", laptop.GetBrand())
		log.Print(" + name: ", laptop.GetName())
		log.Print(" + cpu cores: ", laptop.GetCpu().GetNumberCores())
		log.Print(" + cpu min ghz: ", laptop.GetCpu().GetNumberThreads())
		log.Print(" + ram: ", laptop.GetRam().GetValue(), laptop.GetRam().GetUnit())
		log.Print(" + price: ", laptop.GetPriceUsd())
	}

}

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	// grpc的方式和server建联
	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatal("can not dail server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	// random create 10 laptop
	for i := 0; i < 10; i++ {
		createLaptop(laptopClient)
	}

	// search laptop
	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam: &pb.Memory{
			Value: 8,
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	searchLaptop(laptopClient, filter)
}
