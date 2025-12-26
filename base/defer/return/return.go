package main

import "fmt"

func returnByDefer() (result string) {
	fmt.Println("start defer ... ")
	defer func() {
		fmt.Println("end defer ... ")
		result = "test-defer"
	}()
	fmt.Println("complete defer ... ")
	result = "test-01"
	//return "test-01"
	return result
}

func main() {
	ret := returnByDefer()
	fmt.Println(ret)
}
