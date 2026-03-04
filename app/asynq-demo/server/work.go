// server/main.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mikeian-for-golang/app/asynq-demo/task"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hibiken/asynq"
)

// 任务处理器
func handleEmailDeliveryTask(ctx context.Context, t *asynq.Task) error {
	var p task.EmailDeliveryPayload
	json.Unmarshal(t.Payload(), &p)
	//if err := t.UnmarshalPayload(&p); err != nil {
	//	return err
	//}
	log.Printf("📧 Processing email delivery: %s, time %v", p.String(), time.Now())
	// 模拟耗时操作
	//time.Sleep(2 * time.Second)
	//log.Printf("✅ Email sent to %s", p.Email)
	//fmt.Println("handleEmailDeliveryTask ", p)
	return nil
}

func handleImageResizeTask(ctx context.Context, t *asynq.Task) error {
	//var p task.ImageResizePayload
	//if err := t.UnmarshalPayload(&p); err != nil {
	//	return err
	//}
	//log.Printf("🖼️  Processing image resize: %s", p.String())
	//// 模拟耗时操作
	//time.Sleep(3 * time.Second)
	//log.Printf("✅ Image resized: %s", p.OutputURL)

	fmt.Println("handleImageResizeTask")

	return nil
}

func main() {
	// 1. 创建 Redis 连接
	redisOpt := asynq.RedisClientOpt{Addr: "localhost:6379"}
	// 2. 创建 Asynq Server
	srv := asynq.NewServer(
		redisOpt,
		asynq.Config{
			Concurrency: 10, // 并发 worker 数量
			Queues: map[string]int{
				"critical": 6, // 高优先级队列
				"default":  3, // 默认队列
				"low":      1, // 低优先级队列
			},
			StrictPriority: true,
		},
	)

	// 3. 注册任务处理器
	mux := asynq.NewServeMux()
	mux.HandleFunc(task.TypeEmailDelivery, handleEmailDeliveryTask)
	//mux.HandleFunc(task.TypeImageResize, handleImageResizeTask)

	log.Println("🚀 Asynq server starting...")
	log.Println("📊 Monitor UI: http://localhost:8080")

	// 4. 优雅关闭
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-sigChan
		log.Printf("🛑 Received signal: %v, shutting down...\n", s)
		srv.Shutdown()
	}()

	// 5. 启动服务
	if err := srv.Run(mux); err != nil {
		log.Fatalf("❌ Could not run server: %v", err)
	}
}
