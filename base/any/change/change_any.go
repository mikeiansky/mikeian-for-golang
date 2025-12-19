package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 修改 any 值的内容
func changeAny(data any) any {
	// 类型断言获取 Person
	if p, ok := data.(Person); ok {
		p.Age = p.Age + 1
		p.Name = p.Name + p.Name
		return p // 返回修改后的副本
	}
	return data
}

func main() {
	p := Person{
		Name: "wen",
		Age:  20,
	}
	fmt.Println("原始值:", p)

	// 传递值的副本，函数内修改不会影响原值
	result := changeAny(p)
	fmt.Println("函数返回:", result)
	fmt.Println("原始值未改变:", p)
}
