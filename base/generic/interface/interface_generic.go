package main

import "fmt"

// 定义 Stringer 约束
type Stringer interface {
	String() string
}

func PrintString[T Stringer](value T) {
	fmt.Println(value.String())
}

// 实现自定义类型
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d years old)", p.Name, p.Age)
}

func main() {
	// 使用示例
	person := Person{Name: "Alice", Age: 25}
	PrintString(person) // 输出: Alice (25 years old)
}
