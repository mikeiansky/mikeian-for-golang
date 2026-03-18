package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("time zone app start ... ")
	now := time.Now()

	fmt.Println("time.Now():", now)
	fmt.Println("time.Now().UTC():", now)

	nowStr := now.Format(time.DateTime)
	fmt.Println("nowStr:", nowStr)

	pt, _ := time.Parse(time.DateTime, nowStr)
	fmt.Println("pt:", pt)

	nowUtcStr := now.UTC().Format(time.DateTime)
	fmt.Println("nowUtcStr:", nowUtcStr)

	ptUtc, _ := time.Parse(time.DateTime, nowUtcStr)
	fmt.Println("ptUtc:", ptUtc)

	ptUtcUtc := ptUtc.UTC()
	fmt.Println("ptUtcUtc:", ptUtcUtc)

	fmt.Println("time zone app complete ... ")

}
