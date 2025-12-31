package main

import (
	"fmt"
	"strings"
)

var minorUnits = map[string]int{
	"CNY": 2,
	"USD": 2,
}

func getMinorUnit(currency string) int {
	c := strings.ToUpper(strings.TrimSpace(currency))
	return minorUnits[c]
}

func main() {
	fmt.Println("CNY minor unit", getMinorUnit("CNY"))
	fmt.Println("USD minor unit", getMinorUnit("USD"))
	fmt.Println("JPY minor unit", getMinorUnit("JPY"))
}
