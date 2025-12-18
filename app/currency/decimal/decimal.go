package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {

	ret := decimal.NewFromInt(103).Shift(-2)
	fmt.Println(ret)
	fmt.Println(ret.Value())
	fmt.Println(ret.String())

	var dp *decimal.Decimal
	fmt.Println(dp)
	dp = &ret
	fmt.Println(dp)
	fmt.Println(dp.String())
	fmt.Println(dp.IntPart())
	fmt.Println(dp.Float64())
	fmt.Println(dp.Floor())

	r2 := ret.Mul(*dp)
	fmt.Println(r2)
	fmt.Println(r2.IntPart())
	fmt.Println(r2.RoundBank(2))
	fmt.Println(r2.RoundDown(2))
	fmt.Println(r2.RoundUp(2))
	fmt.Println(r2.RoundFloor(2))
	fmt.Println(r2.RoundCeil(0))
	fmt.Println(r2.RoundFloor(0))

	nd := decimal.NewNullDecimal(decimal.NewFromInt(22))
	fmt.Println(nd)
	fmt.Println(nd.Decimal)

	nd.Valid = false
	fmt.Println(nd.Decimal)
	fmt.Println(nd)
}
