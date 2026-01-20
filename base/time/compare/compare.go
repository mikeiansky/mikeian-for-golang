package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)

	an := now.Add(time.Hour)
	fmt.Println(an)

	fmt.Println(an.Before(now))
	fmt.Println(an.After(now))

}
