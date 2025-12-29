package main

import (
	"fmt"
	"time"
)

func calcSecond(from time.Time, to time.Time) int64 {
	d := to.Sub(from)
	return int64(d.Seconds())
}

func calcMilliseconds(from time.Time, to time.Time) int64 {
	d := to.Sub(from)
	return d.Milliseconds()
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

	fmt.Println("seconds diff : ", calcSecond(now, et))
	fmt.Println("milliseconds diff : ", calcMilliseconds(now, et))

}
