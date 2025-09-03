package main

import "fmt"

type Greet interface {
	SayHello()
}

type Hello struct {
	Msg string
}

func (h *Hello) SayHello() {
	fmt.Println("hello", h.Msg)
}

func main() {
	fmt.Println("app start ... ")

	h := Hello{
		Msg: "world",
	}

	h.SayHello()

	fmt.Println("app complete ...")
}
