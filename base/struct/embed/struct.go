package main

import "fmt"

type Person struct {
	name string
	age  int
	Head struct {
		shape string
		size  int
	}
}

func main() {

	fmt.Println("app start ... ")

	p := Person{
		name: "mikeian",
		age:  23,
		Head: struct {
			shape string
			size  int
		}{shape: "change", size: 20},
	}

	fmt.Println(p)
	fmt.Println(p.Head.shape)
	fmt.Println(p.Head.size)

	fmt.Println("app complete ...")

}
