package main

import (
	"fmt"
	"strconv"
)

func main() {

	str := "124"
	iv, _ := strconv.ParseInt(str, 10, 64)
	//iv := int64(str)
	fmt.Println(iv)

	var ia interface{}
	ia = iv
	fmt.Println(ia)
	//fmt.Println(string(iv))
	fmt.Println(fmt.Sprintf("%v", ia))

}
