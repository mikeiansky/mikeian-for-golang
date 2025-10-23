package main

import "fmt"

func main() {
	fmt.Println("f001 start ...")
	echo("from f001")
	tag := Tag{
		Name: "test from f002",
		Size: 33,
	}
	fmt.Println(tag)
	fmt.Println("f001 complete ...")
}
