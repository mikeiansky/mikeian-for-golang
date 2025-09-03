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

func main() {
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
