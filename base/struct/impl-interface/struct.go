package main

import "fmt"

type Person struct {
	Name string
}

type Greet interface {
	SayHello()
}

func (p *Person) SayHello() {
	fmt.Println("hello i am", p.Name)
}

func UseGreet(g Greet) {
	g.SayHello()
}

func UseGreetPoint(g *Greet) {
	fmt.Println("*g is", *g, ", g is", g, ", &g is", &g)
	(*g).SayHello()
}

func main() {
	fmt.Println("app start ...")

	p := Person{"Bob"}

	p.SayHello()
	UseGreet(&p)
	var g Greet = &p
	fmt.Println("g value is", g)
	var pg = &g
	fmt.Println("pg value is", pg)
	UseGreetPoint(pg)

	fmt.Println("app complete ...")
}
