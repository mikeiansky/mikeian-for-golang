package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/consul/api"
)

func main() {
	// 1. 连接到 Consul 代理 (你的 Docker 容器地址)
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	// 2. 定义服务信息
	serviceID := "my-service-1"
	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "demo-service", // 服务名
		Port:    8010,           // 你的程序监听端口
		Address: "172.17.0.1",   // 你的宿主机 IP (Consul 容器能访问到的)
		Check: &api.AgentServiceCheck{
			HTTP:     "http://172.17.0.1:8010/health",
			Interval: "5s",
			Timeout:  "1s",
		},
	}

	// 3. 注册服务
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("注册失败:", err)
	}
	fmt.Println("服务注册成功！")

	// 4. 启动一个简单的 HTTP Server 供 Consul 检查
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("health!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "I am alive!")
	})

	fmt.Println("正在监听 8010 端口...")
	err = http.ListenAndServe(":8010", nil)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	<-ch
}
