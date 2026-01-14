package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println("start godotenv ... ")
	// 1. 默认加载根目录下的 .env 文件
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. 使用标准库 os.Getenv 读取环境变量
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	apiKey := os.Getenv("API_KEY")

	fmt.Printf("服务运行端口: %s\n", port)
	fmt.Printf("数据库地址: %s\n", dbUrl)
	fmt.Printf("API 密钥: %s\n", apiKey)
	fmt.Println("complete godotenv ... ")

}
