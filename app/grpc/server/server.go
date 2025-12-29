package main

import (
	"context"
	"log"
	"mikeian-for-golang/app/grpc/api"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 实现 Greeter 服务
type server struct {
	api.UnimplementedHelloServiceServer // 必须嵌入，提供默认实现
}

// 实现 SayHello 方法
func (s *server) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())

	// 1. 从 Context 中获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("未找到传入的元数据")
		return nil, status.Error(codes.Unauthenticated, "Metadata required")
	}

	// 2. 读取特定的键值

	// 读取 "auth-token"
	if tokens, found := md["auth-token"]; found && len(tokens) > 0 {
		log.Printf("接收到的 Auth Token: %s", tokens[0])
		// 在此进行身份验证等操作
		if tokens[0] != "bearer-12345" {
			return nil, status.Error(codes.Unauthenticated, "Invalid token")
		}
	} else {
		log.Println("Auth Token 未提供")
	}

	// 读取 "trace-id"
	if traces, found := md["trace-id"]; found {
		log.Printf("接收到的 Trace ID: %s", traces)
	}

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
