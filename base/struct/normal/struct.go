package main

import "fmt"

type Book struct {
	Title string
	Price int
}

func UpdatePointBookTitle(b *Book, nt string) {
	b.Title = nt
}

func UpdateCopyBookTitle(b Book, nt string) {
	b.Title = nt
}

func normal01() {
	fmt.Println("app start ... ")

	b := Book{
		Title: "golang",
		Price: 99,
	}

	fmt.Println("init book", b)

	UpdatePointBookTitle(&b, "python")
	fmt.Println("after update point book", b)

	UpdateCopyBookTitle(b, "javascript")
	fmt.Println("after update copy book", b)

	fmt.Println("app complete ... ")
}

type orderStatus string

type OrderParam struct {
	s orderStatus
}

func normal02() {
	fmt.Println("app start ... ")
	o := OrderParam{}
	fmt.Println("init order param", o.s, "end")
	fmt.Println("after init order param", o.s == "")

}

func main() {
	//normal01()
	normal02()
}
