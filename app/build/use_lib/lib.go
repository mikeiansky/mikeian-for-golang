package main

import (
	"fmt"
	"mikeian-for-golang/app/lib"
)

func main() {

	fmt.Println("normal use lib app start ... ")

	tag := "normal use lib app"
	ret := lib.CreateObj(tag)
	fmt.Println("create obj is", ret)

	fmt.Println("normal use lib app complete ... ")

}
