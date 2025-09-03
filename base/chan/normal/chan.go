package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("app start ...")
	ch := make(chan int)

	go func() {
		time.After(1 * time.Second)
		ch <- 10086
	}()

	ret := <-ch
	fmt.Println("read chan ret", ret)

	fmt.Println("app complete ...")

}
