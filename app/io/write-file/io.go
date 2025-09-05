package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("app start ... ")

	fp := "app/io/write-file/README.md"
	data := []byte("hello world")
	os.WriteFile(fp, data, 0777)

	fmt.Println("app complete ... ")
}
