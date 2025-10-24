package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("test log start ...")

	fmt.Println("test println")
	log.Println("test log 1")
	log.Printf("test log 2")
	// 这个打印致命错误，并且会退出程序
	log.Fatalf("test log 3")
	// 这个不会打印
	log.Fatalf("test log 4")
	fmt.Println("test log complete ...")

}
