package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("app start ...")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("sleeping for 10 seconds, try sending a signal (Ctrl+C or kill) during this time...")
	time.Sleep(10 * time.Second) // 模拟程序运行中，先做些事情，比如初始化

	so := <-ch
	fmt.Println("receive signal ", so)
	fmt.Println("app complete ...")
}
