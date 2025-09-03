package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("app start ... ")

	cli := http.Client{}

	url := "https://www.baidu.com"
	resp, err := cli.Get(url)
	if err != nil {
		fmt.Println("request baidu error", err.Error())
	} else {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read response error", err.Error())
		} else {
			fmt.Println(string(msg))
		}
	}

	fmt.Println("app complete ... ")

}
