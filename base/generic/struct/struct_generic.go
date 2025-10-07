package main

import "fmt"

// 非泛型结构体
//type MyBox struct {
//	Value int
//}
//
//// 为非泛型结构体实现一个方法
//func (b MyBox) Show() {
//	fmt.Println("Non-Generic MyBox:", b.Value)
//}

// 泛型结构体
type MyBox[T any] struct {
	Value T
}

// 为泛型结构体实现一个方法
func (b MyBox[T]) Show() {
	fmt.Println("Generic MyBox:", b.Value)
}

func main() {
	// 使用非泛型版本
	//b1 := MyBox{Value: 42}
	//b1.Show() // 调用非泛型的 Show()

	// 使用泛型版本（T = string）
	b2 := MyBox[string]{Value: "Hello"}
	b2.Show() // 调用泛型版本的 Show()
}
