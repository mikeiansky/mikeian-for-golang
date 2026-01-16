package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

// 定义配置对应的结构体
type AppConfig struct {
	HttpPort int    `json:"http_port"`
	Timeout  string `json:"timeout"`
}

func main() {
	// 1. 初始化 Consul 客户端配置
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500" // 你的 Consul 地址

	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalf("无法创建客户端: %v", err)
	}

	// 2. 从 KV 存储中获取配置
	key := "demo/server/settings"
	pair, _, err := client.KV().Get(key, nil)
	if err != nil {
		log.Fatalf("读取配置出错: %v", err)
	}

	if pair == nil {
		log.Fatalf("未找到配置键: %s", key)
	}

	fmt.Println("pair value is ", string(pair.Value))

	// 3. 将获取到的 Value (字节流) 解析到结构体
	var appConfig AppConfig
	err = json.Unmarshal(pair.Value, &appConfig)
	if err != nil {
		log.Fatalf("解析 JSON 失败: %v", err)
	}

	// 4. 使用配置
	fmt.Printf("配置读取成功!\n端口: %d\n超时时间: %s\n",
		appConfig.HttpPort, appConfig.Timeout)
}
