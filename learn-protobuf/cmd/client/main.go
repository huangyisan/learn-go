package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	pb "learn-protobuf/pb"
	"learn-protobuf/sample"
	"log"
	"os"
	"path/filepath"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func createLaptop(laptopClient pb.LaptopServiceClient, laptop *pb.Laptop) {
	// laptop.Id = ""
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

func uploadImage(laptopClient pb.LaptopServiceClient, laptopID string, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("can not open image file:", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	// 这边启一个stream， 下面代码会使用stream.Send方法发送图片info和多次图片data，用的同一股流，然后在server端能正确的进行对这股流获取info或者组装成data，应该是用到了http2的特性。
	stream, err := laptopClient.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}
	// UploadImageRequest 在proto定义为oneof类型，所以存在两个req，一个是上传图片info，另外一个是上传图片data内容
	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laptopID,
				ImageType: filepath.Ext(imagePath),
			},
		},
	}
	// 将请求发送给server
	err = stream.Send(req)
	if err != nil {
		err2 := stream.RecvMsg(nil)
		log.Fatal("can not send image info:", err, err2)
	}
	// 以大块读取文件内容
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		// 返回字节数和err
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}
		// 多次发送image data
		err = stream.Send(req)
		if err != nil {
			err2 := stream.RecvMsg(nil)
			log.Fatal("cannot send chunk to server: ", err, err2)
		}
	}
	// 关闭并接受请求，此时接受了服务端会返回的信息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot reseive response: ", err)
	}
	log.Printf("image uploaded with id: %s, size: %d", res.GetId(), res.GetSize())
}

func testCreateLaptop(laptopClient pb.LaptopServiceClient) {
	createLaptop(laptopClient, sample.NewLaptop())
}

func testSearchLaptop(laptopClient pb.LaptopServiceClient) {
	// random create 10 laptop
	for i := 0; i < 10; i++ {
		createLaptop(laptopClient, sample.NewLaptop())
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

func testUploadImage(laptopClient pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	createLaptop(laptopClient, laptop)
	uploadImage(laptopClient, laptop.GetId(), "tmp/laptop.jpg")
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
	testUploadImage(laptopClient)

}
