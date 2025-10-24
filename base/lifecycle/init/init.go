package main

import "fmt"

// 会把ext也进行调用
func init() {
	fmt.Println("init ... main")
}

func main() {
	fmt.Println("test init lifecycle")
}
