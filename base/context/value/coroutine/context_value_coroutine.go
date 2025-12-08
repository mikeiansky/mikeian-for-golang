package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: 收到取消信号，停止工作。原因: %v\n", id, ctx.Err())
			return
		default:
			// 模拟从 context 中获取元信息，比如 userID
			idVal := ctx.Value("userID")
			fmt.Printf("Worker %d 正在处理，userID = %v\n", id, idVal)
			time.Sleep(1 * time.Second)
		}
	}
}

type ContextKey string

func main() {
	// 1. 创建一个可取消的 context，并附带一个值（比如 userID）
	var userIdKey ContextKey = "UserID"
	ctx := context.WithValue(context.Background(), userIdKey, 12345)

	// 2. 让这个 c
	//ontext 支持取消
	ctx, cancel := context.WithCancel(ctx)

	// 3. 启动多个 worker，传入同一个 ctx
	go worker(ctx, 1)
	go worker(ctx, 2)

	// 4. 主线程等待 3 秒后，主动取消所有 worker
	time.Sleep(3 * time.Second)
	cancel() // 发送取消信号

	// 5. 给 workers 一点时间响应
	time.Sleep(500 * time.Millisecond)
	fmt.Println("主函数结束")
}
