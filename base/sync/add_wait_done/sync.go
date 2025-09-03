package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("app start ...")
	sc := sync.WaitGroup{}

	ch := make(chan int)

	go func() {
		fmt.Println("sync add 1")
		sc.Add(1)
		close(ch)
	}()

	go func() {
		<-ch
		fmt.Println("sync done")
		sc.Done()
	}()

	<-ch
	sc.Wait()
	fmt.Println("app complete ...")

}
