package main

import (
	"errors"
	"fmt"
)

var e1 = errors.New("e1 ===> ")
var e2 = errors.New("e2 ===> ")

func main() {

	var er error
	if er == nil {
		fmt.Println("hello world")
	} else {
		fmt.Println(er)
	}
	fmt.Println(errors.Is(er, e2))

	er = e1
	e3 := fmt.Errorf("%w: %w", er, e2)
	fmt.Println(e3)
	fmt.Println(errors.Is(e3, e2))
	fmt.Println(errors.Is(e3, e1))

}
