package main

import "fmt"

type IPerson interface {
	SayHello()
}

type Normal struct {
}

func (n Normal) SayHello() {
	fmt.Println("hello normal")
}

func main() {

	n := Normal{}
	n.SayHello()

	var np IPerson = n
	fmt.Println(np)

	var pp IPerson = &n
	fmt.Println(pp)

	nn, ok := np.(Normal)
	fmt.Println(ok, nn)

	pn, ok := pp.(*Normal)
	fmt.Println(ok, pn)
}
