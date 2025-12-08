package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	fmt.Println("my casbin application start ... ")
	e, err := casbin.NewEnforcer("app/casbin/model.conf", "app/casbin/policy.csv")
	fmt.Println("e ", e)
	fmt.Println("err ", err)

	results, err := e.BatchEnforce([][]interface{}{
		{"alice", "data1", "read"},
		{"bob", "data2", "write"},
		{"jack", "data3", "read"},
	})
	if err != nil {
		fmt.Println("err ", err)
	}
	for index, result := range results {
		fmt.Println(index, "result", result)
	}

	fmt.Println("my casbin application complete ... ")

}
