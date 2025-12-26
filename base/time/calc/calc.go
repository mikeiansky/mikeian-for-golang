package main

import (
	"fmt"
	"time"
)

func calculateSecond(t1 time.Time, t2 time.Time) {
	d := t1.Sub(t2)
	fmt.Println("calculate second", d.Seconds())
}

func main() {
	expiredTime := "2025-12-31 22:00:11"
	layout := "2006-01-02 15:04:05"
	et, err := time.Parse(layout, expiredTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(et)

	now := time.Now().UTC()
	fmt.Println(now)

	calculateSecond(now, et)

}
