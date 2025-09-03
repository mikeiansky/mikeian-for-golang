package main

import "fmt"

func UseWriteOnlyChain(wc chan<- int) {
	go func() {
		wc <- 10086
	}()

	// can not read value from chan

	//ret := <- wc
}

func main() {
	fmt.Println("app start ... ")

	ch := make(chan int)

	UseWriteOnlyChain(ch)

	ret := <-ch
	fmt.Println("write only chain ret", ret)

	fmt.Println("app complete ... ")

}
