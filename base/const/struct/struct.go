package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 这样有问题
//const person = Person{
//	Name: "mike",
//	Age:  20,
//}

func main() {
	fmt.Println("const struct start  ... ")
	fmt.Println("const struct complete  ... ")
}
