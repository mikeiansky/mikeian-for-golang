package main

import (
	"fmt"
	"time"
)

func t1() {
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

	fmt.Println("now ", time.Now())
}

func t2() {
	now := time.Now()
	nowUtc := now.UTC()
	fmt.Println(now)
	fmt.Println(nowUtc)
	fmt.Println(now.Unix())
	fmt.Println(nowUtc.Unix())
}

func t3() {
	st := time.UnixMilli(0)
	fmt.Println(st)
	fmt.Println(st.UTC())

	ux := time.Now().UnixMilli()
	at := ux/1000 + 0
	att := time.UnixMilli(at)
	fmt.Println(att)
}

// ParsePaidTime .
func ParsePaidTime(s string) (time.Time, error) {
	ret, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		ret, err = time.Parse(time.DateTime, s)
		if err != nil {
			return time.Time{}, err
		}
	}
	return ret, nil
}

func t4() {
	test := "2026-04-08T11:10:28"
	t, err := ParsePaidTime(test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}

func t5() {
	//paidTime := time.Unix(0/1000, 0).UTC()
	//fmt.Println(paidTime)

	t, err := time.Parse(time.RFC3339, "2026-05-12T08:57:17Z")
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println(t)

	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))

}

func main() {
	//t1()
	//t2()
	//t3()
	// 1775529693248
	// 1775529741014
	// 1775549173425
	//fmt.Println(time.Now().UnixMilli())

	//t4()
	t5()
}
