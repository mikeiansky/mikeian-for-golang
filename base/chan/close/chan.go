package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("app start ...")

	ch := make(chan int)
	sw := sync.WaitGroup{}
	size := 5
	sw.Add(size)

	go func() {
		fmt.Println("read ch start")
		for i := 0; i < size; i++ {
			rc, ok := <-ch
			fmt.Println("read ch i", i, "rc", rc, "ok", ok)
			sw.Done()
		}
		fmt.Println("read ch complete")
	}()

	go func() {
		ch <- 10086
		ch <- 10010
		close(ch)
	}()

	sw.Wait()
	fmt.Println("app complete ...")

}
