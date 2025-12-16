package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("time zone app start ... ")

	fmt.Println("time.Now().UTC():", time.Now().UTC())
	fmt.Println("time.Now():", time.Now())

	fmt.Println("time zone app complete ... ")

}
