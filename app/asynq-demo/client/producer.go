// client/main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"mikeian-for-golang/app/asynq-demo/task"

	"github.com/hibiken/asynq"
)

var (
	redisOpt = asynq.RedisClientOpt{Addr: "localhost:6379"}
)

func main() {
	// 创建客户端
	client := asynq.NewClient(redisOpt)
	defer client.Close()

	log.Println("📤 Asynq client started. Sending tasks...")

	// 发送不同类型的任务
	for i := 0; i < 3; i++ {
		// 随机延迟，模拟不同时间发送
		sendEmailTask(client, i)
	}

	log.Println("✅ All tasks sent! Check Asynqmon at http://localhost:8080")
}

func sendEmailTask(client *asynq.Client, index int) {
	payload := &task.EmailDeliveryPayload{
		UserID:     index + 9080,
		Email:      fmt.Sprintf("user%d@example.com", index),
		TemplateID: "welcome_email",
	}
	pd, _ := json.Marshal(payload)
	task := asynq.NewTask(task.TypeEmailDelivery, pd)

	// 发送到 critical 队列，优先级最高
	//info, err := client.Enqueue(task, asynq.Queue("critical"), asynq.MaxRetry(3))

	//tt := time.Now().Add(time.Second * 5)
	//pa := asynq.ProcessAt(tt)
	//fmt.Printf("task %v , at time : %v \n", task, tt)

	//info, err := client.EnqueueContext(context.Background(), task, pa)
	//info, err := client.Enqueue(task, asynq.Queue("default"))
	info, err := client.Enqueue(task, asynq.Queue("critical"))
	if err != nil {
		log.Printf("❌ Failed to enqueue email task: %v", err)
		return
	}

	log.Printf("📧 Enqueued email task: ID=%s Queue=%s", info.ID, info.Queue)

}

//
//func sendImageTask(client *asynq.Client, index int) {
//	payload := &task.ImageResizePayload{
//		ImageURL:  fmt.Sprintf("https://example.com/images/%d.jpg", index),
//		Width:     800,
//		Height:    600,
//		OutputURL: fmt.Sprintf("https://cdn.example.com/resized/%d.jpg", index),
//	}
//
//	task := asynq.NewTask(task.TypeImageResize, payload)
//
//	// 发送到 default 队列
//	info, err := client.Enqueue(task, asynq.Queue("default"))
//	if err != nil {
//		log.Printf("❌ Failed to enqueue image task: %v", err)
//		return
//	}
//
//	log.Printf("🖼️  Enqueued image task: ID=%s Queue=%s", info.ID, info.Queue)
//}
//
//func sendLowPriorityEmailTask(client *asynq.Client, index int) {
//	payload := &task.EmailDeliveryPayload{
//		UserID:     index + 2000,
//		Email:      fmt.Sprintf("newsletter%d@example.com", index),
//		TemplateID: "newsletter_weekly",
//	}
//
//	task := asynq.NewTask(task.TypeEmailDelivery, payload)
//
//	// 发送到 low 队列，优先级最低
//	info, err := client.Enqueue(task, asynq.Queue("low"), asynq.ProcessIn(5*time.Second))
//	if err != nil {
//		log.Printf("❌ Failed to enqueue low priority task: %v", err)
//		return
//	}
//
//	log.Printf("📧 Enqueued low priority task: ID=%s Queue=%s", info.ID, info.Queue)
//}
