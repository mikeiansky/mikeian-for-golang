package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("app start ...")
	ch := make(chan int, 0)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("before send")
		ch <- 10086
		fmt.Println("after send, and send again")
		ch <- 10087
		fmt.Println("after twice send")

	}()

	fmt.Println("ready read ch")
	time.Sleep(3 * time.Second)
	fmt.Println("before read ch")
	ret := <-ch
	fmt.Println("read chan ret", ret)
	time.Sleep(1 * time.Second)
	fmt.Println("app complete ...")

}
