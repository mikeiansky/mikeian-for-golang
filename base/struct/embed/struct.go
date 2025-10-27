package main

import "fmt"

type Person struct {
	name string
	age  int
	H1   struct {
		shape string
		size  int
	}
}

type Head2 struct {
	shape string
	size  int
}

type Person2 struct {
	name string
	age  int
	H2   Head2
}

func main() {

	fmt.Println("app start ... ")

	p := Person{
		name: "mikeian",
		age:  23,
		H1: struct {
			shape string
			size  int
		}{shape: "change", size: 20},
	}

	fmt.Println(p)
	fmt.Println(p.H1.shape)
	fmt.Println(p.H1.size)

	h2 := Head2{
		shape: "pointer",
		size:  20,
	}

	p2 := Person2{
		name: "mikeian",
		age:  23,
		H2:   h2,
	}

	fmt.Println(p2)
	fmt.Println(p2.H2.shape)
	fmt.Println(p2.H2.size)

	fmt.Println("app complete ...")

}
