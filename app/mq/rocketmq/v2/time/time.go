package main

import (
	"fmt"
	"strconv"
	"time"
)

// calculateDeliveryTimestampMs 计算指定延迟时间后的毫秒级时间戳字符串
func calculateDeliveryTimestampMs(delay time.Duration) string {
	// 1. 确定目标投递时间点
	deliveryTime := time.Now().Add(delay)

	// 2. 转换为毫秒级 Unix 时间戳
	// deliveryTime.UnixNano() / 1e6 相当于 time.UnixMilli()
	deliveryTimestampMs := deliveryTime.UnixNano() / int64(time.Millisecond)

	// 3. 转换为字符串格式，这是 RocketMQ 属性的要求
	return strconv.FormatInt(deliveryTimestampMs, 10)
}

func main() {
	ret := calculateDeliveryTimestampMs(time.Second * 48)
	fmt.Println(ret)
}
