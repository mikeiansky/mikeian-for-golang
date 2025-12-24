package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

func sendSyncMessage(p rocketmq.Producer, message *primitive.Message) (*primitive.SendResult, error) {
	res, err := p.SendSync(context.Background(), message)
	return res, err
}

func main() {
	p, _ := rocketmq.NewProducer(
		producer.WithNameServer([]string{"127.0.0.1:9876"}),
		producer.WithRetry(2),
	)

	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		return
	}
	defer func(p rocketmq.Producer) {
		err := p.Shutdown()
		if err != nil {
			fmt.Printf("shutdown producer error: %s", err.Error())
		}
	}(p)

	msg := &primitive.Message{
		Topic: "TestTopic",
		Body:  []byte("Hello RocketMQ Go Client, this is sync message 6!"),
	}
	res, err := sendSyncMessage(p, msg)

	if err != nil {
		fmt.Printf("send sync message error: %s\n", err)
	} else {
		fmt.Printf("send sync message success: result=%s\n", res.String())
	}

}
