package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("app start ... ")

	fp := "app/io/read-file/README.md"
	data, err := os.ReadFile(fp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	fmt.Println("app complete ... ")
}
