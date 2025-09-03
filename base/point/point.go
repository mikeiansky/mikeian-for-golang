package main

import "fmt"

func SwapValue(a *int, b *int) {
	*a, *b = *b, *a
	fmt.Println("a value is", a, ",*a value is", *a, ",&a value is", &a)
}

func main() {

	fmt.Println("app start ... ")
	a, b := 3, 4
	fmt.Println("before swap a", a, "b", b)
	SwapValue(&a, &b)
	fmt.Println("after swap a", a, "b", b)
	fmt.Println("app complete ... ")

}
