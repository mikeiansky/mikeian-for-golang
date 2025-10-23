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

func useInterface(it interface{}) {
	fmt.Println("use interface", it)
}

func main() {
	fmt.Println("app start ... ")

	h := Hello{
		Msg: "world",
	}

	h.SayHello()

	useInterface(h)

	fmt.Println("app complete ...")
}
