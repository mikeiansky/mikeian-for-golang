package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// chapter3/sources/evaluation_order_1.go
var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func f() int {
	d++
	return d
}

func hhh() {
	sw := sync.WaitGroup{}
	sw.Add(1)
	var ne error
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("after")
		sw.Done()
		ne = errors.New("time out")
	})
	sw.Wait()
	//fmt.Println("after func ")
	fmt.Println("after func ", ne.Error())
}

func main() {
	fmt.Println(a, b, c, d)
	errors.New("test")
	ret := fmt.Sprintf("%d", time.Now().UnixMilli())
	fmt.Println(ret)
	hhh()
	time.Sleep(2 * time.Second)
	fmt.Println("complete ")
}
