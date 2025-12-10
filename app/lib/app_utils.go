package lib

import "fmt"

func CreateObj(p any) any {
	fmt.Println("create obj start ... param is", p)
	return p
}
