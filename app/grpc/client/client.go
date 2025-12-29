package main

import (
	"context"
	"log"
	"mikeian-for-golang/app/grpc/api"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 连接到 gRPC 服务端
	//conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// other server

	conn, err2 := grpc.NewClient("172.27.116.220:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err2 != nil {
		panic(err2)
	}
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	c := api.NewHelloServiceClient(conn)

	// 从命令行参数获取 name，如果没有则默认为 "World"
	name := "World 2"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	// 调用 SayHello 方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 1. 创建元数据
	md := metadata.New(map[string]string{
		"auth-token": "bearer-12345", // 注意：键必须是小写
		"trace-id":   "abcdefg",
	})
	ctx = metadata.NewOutgoingContext(ctx, md)
	r, err := c.SayHello(ctx, &api.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("🟢 Response from server: %s", r.GetMessage())
}
