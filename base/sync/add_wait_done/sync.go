package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("app start ...")
	sc := sync.WaitGroup{}

	ch := make(chan int)

	go func() {
		fmt.Println("sync add 1")
		sc.Add(1)
		//close(ch)
		time.Sleep(2 * time.Second)
		ch <- 1
	}()
	fmt.Println("after go 1")

	go func() {
		<-ch
		fmt.Println("sync done")
		//sc.Add(-1)
		sc.Done()
	}()
	fmt.Println("after go 2")
	time.Sleep(1 * time.Second)
	//<-ch
	sc.Wait()
	fmt.Println("after sc wait")
	fmt.Println("app complete ...")

}
