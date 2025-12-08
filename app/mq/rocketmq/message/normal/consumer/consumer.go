package main

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithNameServer([]string{"127.0.0.1:9876"}),
		consumer.WithGroupName("testGroup"),
	)

	err := c.Subscribe("TestTopic", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			fmt.Printf("Received message: %s\n", msg.Body)
		}
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		fmt.Printf("subscribe error: %s\n", err)
		return
	}

	err = c.Start()
	if err != nil {
		fmt.Printf("start consumer error: %s\n", err)
		return
	}
	defer func(c rocketmq.PushConsumer) {
		err := c.Shutdown()
		if err != nil {
			panic(err)
		}
	}(c)

	// 保持程序运行
	select {}
}
