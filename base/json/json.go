package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("app start ... ")

	data := `
		{
			"name": "mikeian",
			"age": 27
		}
	`
	fmt.Println(data)

	person := Person{}
	err := json.Unmarshal([]byte(data), &person)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(person)

	jd, err := json.Marshal(person)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jd))
	}

	fmt.Println("app complete ...")
}
