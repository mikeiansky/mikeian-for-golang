package main

import (
	"fmt"
	"mikeian-for-golang/app/protobuf/person"
)

func main() {
	fmt.Println("app start ... ")

	p := person.Person{
		Name: "mikeian",
		Age:  33,
	}
	fmt.Println(p)

	fmt.Println("app complete ... ")
}
