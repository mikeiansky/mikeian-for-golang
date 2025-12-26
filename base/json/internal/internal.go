package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	p := &person{
		Name: "John Doe",
		Age:  42,
	}
	fmt.Println("normal person %v", p)

	ret, _ := json.Marshal(p)
	fmt.Println(string(ret))

	up := &person{}
	json.Unmarshal(ret, up)
	fmt.Println("unmarshal person %v", up)
	fmt.Println(up.Name)
	fmt.Println(up.Age)

}
