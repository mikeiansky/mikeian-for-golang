package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("app start ... ")

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		msg := r.URL.Query().Get("msg")
		fmt.Println("received msg: ", msg)

		echo := "hello " + msg
		_, err := w.Write([]byte(echo))
		if err != nil {
			fmt.Println("error writing response")
		}

	})

	err := http.ListenAndServe(":20001", nil)
	if err != nil {
		fmt.Println("error starting server, err", err.Error())
	}

	fmt.Println("app complete ... ")
}
