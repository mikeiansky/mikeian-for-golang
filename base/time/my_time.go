package main

import (
	"fmt"
	"time"
)

func main() {
	// 同一个 Unix 时间戳，不同时区显示
	unixTime := int64(1724156611)

	// UTC 时间
	utcTime := time.Unix(unixTime, 0).UTC()
	fmt.Printf("UTC 时间:    %s\n", utcTime.Format("2006-01-02T15:04:05.000Z"))

	// 对应 -07:00 时区的本地时间
	location, _ := time.LoadLocation("America/Denver") // MST = UTC-7
	localTime := time.Unix(unixTime, 0).In(location)
	fmt.Printf("-07:00 时区: %s\n", localTime.Format("2006-01-02T15:04:05.000-07:00"))

	// 对应 +08:00 时区（中国）
	chinaLocation, _ := time.LoadLocation("Asia/Shanghai")
	chinaTime := time.Unix(unixTime, 0).In(chinaLocation)
	fmt.Printf("+08:00 时区: %s\n", chinaTime.Format("2006-01-02T15:04:05.000+08:00"))
}
