package main

import "fmt"

func main() {

	var size int64 = 78
	fmt.Println(size)
	fmt.Println(&size)

	var ip *int64
	fmt.Println(ip)
	ip = new(int64)
	fmt.Println(ip)
	fmt.Println(*ip)
	ip = &size
	fmt.Println(ip)
	fmt.Println(*ip)

	var ri *int64
	fmt.Println(ri)
	// 这里会报错
	//fmt.Println(*ri)

}
