package main

import (
	"fmt"
	"sync"
	"time"
)

func goWait() {
	fmt.Println("go wait start ...")

	sc := sync.WaitGroup{}

	sc.Add(1)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("go sync done start")
		sc.Done()
		fmt.Println("go sync done complete")
	}()

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("go sync wait start")
		sc.Wait()
		fmt.Println("go sync done complete")
	}()

	// 如果这里不等待，会直接完成该方法
	//time.Sleep(3 * time.Second)
	fmt.Println("go wait complete ...")
}

func main() {
	fmt.Println("app start ...")
	goWait()
	fmt.Println("app complete ...")
}
