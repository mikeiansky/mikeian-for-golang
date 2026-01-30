package main

import (
	"fmt"
	"math"
)

func main() {

	n1 := 6413145
	unit := 3
	n2 := math.Pow(10, float64(unit))
	n3 := float64(n1) / n2

	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

}
