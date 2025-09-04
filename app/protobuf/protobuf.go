package main

import (
	"fmt"
	"mikeian-for-golang/app/protobuf/city"
	"mikeian-for-golang/app/protobuf/person"
)

func main() {

	fmt.Println("app start ... ")

	c := city.City{
		Name:    "shenzhen",
		Address: "guangdongshen nanshanqu",
	}

	fmt.Println(c)

	p := person.Person{
		Name: "mikeian",
		Age:  20,
		City: &c,
	}
	fmt.Println(p)

	fmt.Println("app complete ... ")

}
