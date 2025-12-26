package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	fmt.Println("app start ...")
	// ✅ 1. 创建 Redis 客户端（连接到本地 Redis 服务）
	// 默认 Redis 地址是 localhost:6379，无密码，无 DB 选择
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.31.109:6379", // Redis 服务器地址
		Password: "",                    // 没有密码，默认为空
		DB:       0,                     // 默认使用 DB 0
	})

	// ✅ 2. 使用一个 context（通常用 context.Background()）
	ctx := context.Background()

	// ✅ 3. 测试连接：Ping Redis 服务，检查是否正常
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("无法连接到 Redis: %v", err)
	}
	fmt.Println("✅ Redis 连接成功:", pong)

	//// ✅ 4. 写入数据到 Redis（SET key value）
	key := "mykey"
	value := "hello, Redis from Go!"

	err = rdb.Set(ctx, key, value, 10*time.Minute).Err() // 设置 10 分钟后过期
	if err != nil {
		log.Fatalf("写入 Redis 失败: %v", err)
	}
	fmt.Printf("✅ 已写入 Redis - Key: %s, Value: %s (10分钟后过期)\n", key, value)

	// ✅ 5. 从 Redis 读取数据（GET key）
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Println("❌ Redis 中没有找到这个 key:", key)
	} else if err != nil {
		log.Fatalf("从 Redis 读取失败: %v", err)
	} else {
		fmt.Printf("✅ 从 Redis 读取成功 - Key: %s, Value: %s\n", key, val)
	}

	// ✅ 6. 尝试读取一个不存在的 key（演示错误处理）
	nonExistentKey := "non_existent_key"
	val, err = rdb.Get(ctx, nonExistentKey).Result()
	if err == redis.Nil {
		fmt.Printf("ℹ️ Redis 中没有找到 key: %s （这是正常的，演示 Not Found 处理）\n", nonExistentKey)
	} else if err != nil {
		log.Fatalf("读取 key %s 时出错: %v", nonExistentKey, err)
	} else {
		fmt.Printf("✅ 读取到 key %s 的值: %s\n", nonExistentKey, val)
	}
	fmt.Println("app complete ...")
}
