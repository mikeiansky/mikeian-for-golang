package main

import (
	"fmt"
	"time"
)

func CalcMilliseconds(from time.Time, to time.Time) int64 {
	d := to.Sub(from)
	return d.Milliseconds()
}

func main() {
	now := time.Now().UTC()
	delayMilliseconds := CalcMilliseconds(now, now.Add(time.Hour*24))
	delayTime := time.Now().Add(time.Millisecond * time.Duration(delayMilliseconds))
	fmt.Println("delay time : ", delayTime)
}
