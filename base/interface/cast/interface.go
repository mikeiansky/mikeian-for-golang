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

//func UseGreet(g Greet) {
//	g.SayHello()
//}

func UseInterface(g interface{}) {
	h, ok := g.(Greet)
	if ok {
		h.SayHello()
	} else {
		fmt.Println("g is not a Greet")
	}
}

func main() {
	fmt.Println("app start ... ")

	h := Hello{
		Msg: "world",
	}
	// is not a greet interface
	UseInterface(h)
	// is a greet interface
	UseInterface(&h)
	fmt.Println("app complete ... ")
}
