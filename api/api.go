package main

import "fmt"
import "mikeian-for-golang/service"

func main() {

	fmt.Println("api app start ... ")
	s := &service.Service{}
	app := s.CreateApp("api")
	fmt.Println("app create result", app)
	fmt.Println("api app complete ... ")

}
