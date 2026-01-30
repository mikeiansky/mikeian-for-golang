package main

import "fmt"

func main() {

	status := "paied"

	switch status {
	case "paied":
		//fmt.Println("paied")
	case "failed":
		fmt.Println("failed")
	case "complete":
		fmt.Println("complete")
	}

}
