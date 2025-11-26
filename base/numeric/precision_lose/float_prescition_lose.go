package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 经典的浮点数精度问题
	a := 0.1
	b := 0.2
	c := a + b

	d := 1129.6

	var bf big.Float
	bf.SetFloat64(a)

	fmt.Println("a value is ", a)
	fmt.Println("a value is ", bf.String())
	fmt.Println("d * 100 = ", d*100) //输出：112959.99999999999
	fmt.Printf("%.2f\n", 19.90)
	fmt.Println("c = ", c)
	fmt.Printf("0.1 + 0.2 = %.20f\n", c)            // 0.30000000000000004441
	fmt.Printf("0.1 + 0.2 == 0.3 ? %t\n", c == 0.3) // false !!!

	// 金融计算中的灾难
	price1 := 12.34
	price2 := 56.78
	total := price1 + price2

	fmt.Println("total = ", total)
	fmt.Printf("12.34 + 56.78 = %.20f\n", total) // 69.119999999999995
	fmt.Printf("Expected: 69.12, Actual: %.20f\n", total)

	// 累积误差
	var sum float64 = 0
	for i := 0; i < 10; i++ {
		sum += 0.1
	}
	fmt.Println("sum = ", sum)            // 0.99999999999999988898
	fmt.Printf("0.1 * 10 = %.20f\n", sum) // 0.99999999999999988898
	fmt.Printf("Expected: 1.0, Actual: %.20f\n", sum)
}
