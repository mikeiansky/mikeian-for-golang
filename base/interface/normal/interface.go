package main

import "fmt"

type Greet interface {
	SayHello()
	CreateTag(string) string
	NoUseParam(string) string
}

type Hello struct {
	Msg string
}

type Body struct{}

func (h *Hello) SayHello() {
	fmt.Println("hello", h.Msg)
}

func (h *Hello) CreateTag(tag string) string {
	fmt.Println("create tag", tag)
	return tag + " created"
}

func (h *Hello) NoUseParam(string) *Body {
	fmt.Println("no use param")
	return nil
}

type Person struct{}

func CreatePerson() *Person {
	return nil
}

func CreateNonePerson() Person {
	// can not return nill
	return Person{}
}

func useInterface(it interface{}) {
	fmt.Println("use interface", it)
}

func main() {
	fmt.Println("app start ... ")

	h := Hello{
		Msg: "world",
	}

	h.SayHello()
	useInterface(h)

	tag := h.CreateTag("hello")
	fmt.Println("create tag result", tag)

	result := h.NoUseParam(tag)
	fmt.Println("no use param", result)

	person := CreatePerson()
	fmt.Println("create person result", person)

	nonePerson := CreateNonePerson()
	fmt.Println("create none person result", nonePerson)

	fmt.Println("app complete ...")
}
