package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("app start ... ")

	fp := "app/io/write-file/README.md"
	data := []byte("hello world")
	err := os.WriteFile(fp, data, 0777)
	if err != nil {
		return
	}

	fmt.Println("app complete ... ")
}
