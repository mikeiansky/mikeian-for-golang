package main

import (
	"fmt"
	"time"
)

func main() {
	timeStr := "2024-01-15 10:30:45"
	layout := "2006-01-02 15:04:05"

	// 1. 加载北京时间（Asia/Shanghai）的时区信息
	cstLocation, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		// 如果系统找不到时区数据库，通常会返回 nil
		fmt.Println("Error loading location:", err)
		return
	}

	// 2. 使用 ParseInLocation 将字符串解析为 CST 时间
	// 明确告诉 Go：这个字符串表示的是 cstLocation 时区的时间。
	cstTime, err := time.ParseInLocation(layout, timeStr, cstLocation)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	// 3. 将这个时间点（cstTime）转换到 UTC 时区
	utcTime := cstTime.In(time.UTC)

	fmt.Println("--- 原始信息 ---")
	fmt.Printf("原始字符串: %s\n", timeStr)

	fmt.Println("\n--- 解析为北京时间 (CST) ---")
	fmt.Printf("CST 时间: %s\n", cstTime.Format(layout)) // 10:30:45
	fmt.Printf("CST Location: %s\n", cstTime.Location()) // Asia/Shanghai (UTC+8)

	fmt.Println("\n--- 转换为 UTC 时间 ---")
	fmt.Printf("UTC 时间: %s\n", utcTime.Format(layout)) // 02:30:45 (比 CST 少 8 小时)
	fmt.Printf("UTC Location: %s\n", utcTime.Location()) // UTC (UTC+0)
}
