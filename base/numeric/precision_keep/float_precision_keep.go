package main

import (
	"fmt"
	"math/big"
)

func main() {
	// ✅ 使用整数构造，完全避免小数精度问题
	a := new(big.Float).SetPrec(64).SetInt64(1) // 0.1 = 1/10
	// ✅ 查看实际类型
	fmt.Printf("a 的类型: %T\n", a)         // *math/big.Float
	fmt.Printf("a 的值: %s\n", a.String()) // 1

	b := new(big.Float).SetPrec(64).SetInt64(10)
	a.Quo(a, b) // a = 1/10 = 0.1
	fmt.Println("a value is ", a)

	bVal := new(big.Float).SetPrec(64).SetInt64(2)
	bVal.Quo(bVal, big.NewFloat(10)) // b = 2/10 = 0.2
	fmt.Println("bVal value is ", bVal)

	c := new(big.Float).Add(a, bVal)
	expected := new(big.Float).SetPrec(64).SetInt64(3)
	expected.Quo(expected, big.NewFloat(10)) // expected = 3/10 = 0.3

	fmt.Printf("0.1 + 0.2 = %s\n", c.String())
	fmt.Printf("0.1 + 0.2 == 0.3 ? %t\n", c.Cmp(expected) == 0) // true
}
