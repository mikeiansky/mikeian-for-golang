package main

import (
	"fmt"
	"math/big"
)

func convertFloat(x float64) *big.Float {
	sx := big.NewFloat(x)
	fmt.Printf("步骤1 - big.NewFloat(%f) = %s\n", x, sx.String())

	fx, _ := new(big.Float).SetPrec(64).SetString(sx.String())
	fmt.Printf("步骤2 - SetString 结果 = %s\n", fx.String())

	return fx
}

func main() {
	a := 0.1
	b := 0.2

	fmt.Printf("原始 float64:\n")
	fmt.Printf("a = %.20f\n", a)
	fmt.Printf("b = %.20f\n", b)

	fmt.Printf("\n转换过程:\n")
	ba := convertFloat(a)
	bb := convertFloat(b)

	fmt.Printf("\n转换结果:\n")
	fmt.Printf("ba = %s\n", ba.String())
	fmt.Printf("bb = %s\n", bb.String())

	ar := new(big.Float).Add(ba, bb)
	fmt.Printf("\n加法结果:\n")
	fmt.Printf("ar = %s\n", ar.String())
	fmt.Printf("ar Text('f', 20) = %s\n", ar.Text('f', 10))

	// 与精确的 0.3 比较
	expected, _ := new(big.Float).SetPrec(64).SetString("0.3")
	fmt.Printf("\n比较:\n")
	fmt.Printf("期望的 0.3 = %s\n", expected.String())
	fmt.Printf("ar == 0.3 ? %t\n", ar.Cmp(expected) == 0)

	// 显示差值
	diff := new(big.Float).Sub(ar, expected)
	fmt.Printf("差值 = %s\n", diff.Text('e', 10))
}
