package main

import "fmt"

// chapter3/sources/evaluation_order_1.go
var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func f() int {
	d++
	return d
}

func main() {
	fmt.Println(a, b, c, d)
}
