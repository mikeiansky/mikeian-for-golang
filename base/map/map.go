package main

import (
	"encoding/json"
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

	header := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"Test":         "oo2",
	}
	fmt.Println(header)
	ret, _ := json.Marshal(header)
	fmt.Println(string(ret))

	header2 := map[string]string{}
	json.Unmarshal(ret, &header2)
	fmt.Println(header2)

}
