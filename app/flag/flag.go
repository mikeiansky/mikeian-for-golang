package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("app start ... ")
	var flagconf string
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.Parse()
	fmt.Println("flagconf ", flagconf)
	fmt.Println("app complete ... ")
}
