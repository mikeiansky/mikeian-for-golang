package main

import (
	"context"
	"log"
	"mikeian-for-golang/app/grpc/api"
	"net"

	"google.golang.org/grpc"
)

// 实现 Greeter 服务
type server struct {
	api.UnimplementedHelloServiceServer // 必须嵌入，提供默认实现
}

// 实现 SayHello 方法
func (s *server) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &api.HelloReply{Message: "Hello, " + req.GetName() + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterHelloServiceServer(s, &server{})

	log.Println("🟢 gRPC Server is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
