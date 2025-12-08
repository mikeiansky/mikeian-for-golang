package main

import "fmt"

func main() {
	fmt.Println("test lint start ... ")
	// nolint
	analyie := "analyze"
	fmt.Println(analyie)

	fmt.Println("test lint complete ... ")
}
