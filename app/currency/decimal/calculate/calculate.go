package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func CalculateAllFee(amount, fixedFee, ratePercent decimal.NullDecimal) decimal.NullDecimal {
	result := decimal.NullDecimal{}
	if !amount.Valid {
		return result
	}
	result.Valid = true
	if fixedFee.Valid {
		result.Decimal = fixedFee.Decimal
	}
	if ratePercent.Valid {
		result.Decimal = result.Decimal.Add(amount.Decimal.Mul(ratePercent.Decimal.Shift(-2)))
	}
	return result
}

func main() {
	amount := decimal.NewNullDecimal(decimal.NewFromInt(102).Shift(-2))
	fmt.Println(amount)

	fixedFee := decimal.NewNullDecimal(decimal.NewFromInt(11))
	fmt.Println(fixedFee)

	ratePercent := decimal.NewNullDecimal(decimal.NewFromInt(23))
	fmt.Println(ratePercent)

	result := CalculateAllFee(amount, fixedFee, ratePercent)
	fmt.Println(result)

	// result 保留两位小数
	fmt.Println(result.Decimal.Round(2))
	fmt.Println(result.Decimal.Shift(2).Round(0))

}
