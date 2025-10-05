package main

import (
	"context"
	"log"
	"mikeian-for-golang/app/grpc/api"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接到 gRPC 服务端
	//conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// other server
	conn, err := grpc.NewClient("192.168.8.151:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	defer conn.Close()

	c := api.NewHelloServiceClient(conn)

	// 从命令行参数获取 name，如果没有则默认为 "World"
	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// 调用 SayHello 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &api.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("🟢 Response from server: %s", r.GetMessage())
}
