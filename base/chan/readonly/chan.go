package main

import "fmt"

func UseReadonlyChain(rc <-chan int) {

	ret := <-rc

	fmt.Println("read readonly chain ret", ret)

	// can not write value to chan
	//rc <- 10086

}

func main() {
	fmt.Println("app start ...")

	ch := make(chan int)

	go func() {
		ch <- 10086
	}()

	UseReadonlyChain(ch)

	fmt.Println("app complete ...")

}
