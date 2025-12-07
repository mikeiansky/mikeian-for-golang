package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("panic")
	}()

	fmt.Println("app start ...")

	a := make([]int, 3, 5)
	fmt.Println(a, len(a), cap(a))
	//a[3] = 2 // 这样会报错
	a = append(a, 1, 2, 3, 4, 5) // 这里用回增长，且不会受限于capacity
	fmt.Println(a, len(a), cap(a))
}
