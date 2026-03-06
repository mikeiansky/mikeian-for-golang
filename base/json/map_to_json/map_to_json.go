package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	m := map[string]interface{}{
		"name":  "Alice",
		"age":   25,
		"email": "alice@example.com",
	}

	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("转换失败:", err)
		return
	}

	fmt.Println(string(data))

}
