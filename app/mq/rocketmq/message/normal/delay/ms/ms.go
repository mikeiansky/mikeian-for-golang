/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

// calculateDeliveryTimestampMs 计算指定延迟时间后的毫秒级时间戳
func calculateDeliveryTimestampMs(delay time.Duration) string {
	// 目标投递时间点
	deliveryTime := time.Now().Add(delay)

	// 转换为毫秒级 Unix 时间戳 (Go time.UnixNano / 1e6)
	deliveryTimestampMs := deliveryTime.UnixNano() / 1e6

	// 转换为字符串格式，这是 RocketMQ 属性的要求
	return strconv.FormatInt(deliveryTimestampMs, 10)
}

func main() {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
	)
	err := p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s", err.Error())
		os.Exit(1)
	}
	for i := 0; i < 10; i++ {

		// 1. 定义自定义延迟时长（例如：15 分钟 30 秒）
		//customDelay := 15*time.Minute + 30*time.Second
		customDelay := 36 * time.Second

		// 2. 计算目标投递时间戳（毫秒字符串）
		deliveryTimeMsStr := calculateDeliveryTimestampMs(customDelay)

		msg := primitive.NewMessage("TestTopic", []byte(fmt.Sprintf("Hello RocketMQ Go Client! this is delay v2 message , : %d", i)))
		//msg.WithDelayTimeLevel(3)
		msg.WithProperty("__STARTDELIVERTIME", deliveryTimeMsStr)
		res, err := p.SendSync(context.Background(), msg)

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s", err.Error())
	}
}
