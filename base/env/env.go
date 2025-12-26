package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("TEST", "hahah")
	level := os.Getenv("TEST")
	fmt.Println(level)

}
