package main

import (
	"fmt"
	"strconv"
)

func main() {

	id := int64(18723)
	ret := strconv.FormatInt(id, 10)

	fmt.Println(ret)
}
