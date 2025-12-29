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

// (这里包含上面定义的 calculateDeliveryTimestampMs 函数)

func main() {
	topic := "TestTopic"
	// 确保这个地址指向你的 RocketMQ 5.x NameServer
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"127.0.0.1:9876"})),
		producer.WithRetry(2),
	)

	if err := p.Start(); err != nil {
		fmt.Printf("启动生产者失败: %s\n", err.Error())
		os.Exit(1)
	}
	defer p.Shutdown()

	// ---- 自定义延迟逻辑 ----

	// 1. 定义自定义延迟时长 (例如：45 秒)
	customDelay := 45 * time.Second

	// 2. 计算目标投递时间戳（毫秒字符串）
	deliveryTimeMsStr := calculateDeliveryTimestampMs(customDelay)

	// 打印预期时间，便于调试
	expectedDeliveryTime := time.Now().Add(customDelay)
	fmt.Printf("消息将在 %s 投递 (延迟 %v)\n", expectedDeliveryTime.Format("15:04:05.000"), customDelay)

	// 3. 创建消息
	msg := primitive.NewMessage(topic, []byte(fmt.Sprintf("延迟消息内容: %s", expectedDeliveryTime.String())))

	// 4. 【核心步骤】设置自定义延迟属性
	// RocketMQ 5.x Broker 会根据这个绝对时间戳将消息放入时间轮进行调度。
	msg.WithProperty("__STARTDELIVERTIME", deliveryTimeMsStr)

	// 5. 发送消息
	res, err := p.SendSync(context.Background(), msg)

	if err != nil {
		fmt.Printf("发送延迟消息失败: %s\n", err)
	} else {
		fmt.Printf("发送成功: MsgId=%s, Status=%d\n", res.MsgID, res.Status)
	}
}

// --------------------------------------------------------------------------------
// ** 辅助函数 **
func calculateDeliveryTimestampMs(delay time.Duration) string {
	deliveryTime := time.Now().Add(delay)
	deliveryTimestampMs := deliveryTime.UnixNano() / int64(time.Millisecond)
	return strconv.FormatInt(deliveryTimestampMs, 10)
}
