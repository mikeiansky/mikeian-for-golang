package main

import (
	"fmt"
	"time"
)

func main() {

	ret := time.Unix(1724156611, 1235).Format("2006-01-02T15:04:05.000-07:00")
	fmt.Println(ret)

}
