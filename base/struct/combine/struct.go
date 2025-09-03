package main

import "fmt"

type Head struct {
	Shape string
	Size  int
}

type Body struct {
	Shape string
	Width int
}

type Feet struct {
	Height int
	Type   string
}

type Person struct {
	Head
	Body
	Feet
	Name string
	Age  int
}

func (h *Head) Think() {
	fmt.Println(h.Shape, " head is thinking ... and size is ", h.Size)
}

func (b *Body) Eat() {
	fmt.Println(b.Shape, " Body is eat ... and width is ", b.Width)
}

func (f *Feet) Running() {
	fmt.Println(f.Type, " Feet Running and height is ", f.Height)
}

func (p *Person) Live() {
	fmt.Println(p, " Live in the world")
}

func main() {
	fmt.Println("app start ...")

	//h := Head{
	//	Shape: "rect",
	//	Size:  23,
	//}

	p := Person{
		Head: Head{
			Shape: "circle",
			Size:  20,
		},
		Body: Body{
			Shape: "rectangle",
			Width: 101,
		},
		Feet: Feet{
			Type:   "line",
			Height: 99,
		},
		Name: "mike ian",
		Age:  21,
	}

	p.Think()
	p.Eat()
	p.Running()
	p.Live()
	fmt.Println("person Head.shape is", p.Head.Shape)
	fmt.Println("person Body.shape is", p.Body.Shape)
	fmt.Println("person size is", p.Size)
	fmt.Println("person width is", p.Width)
	fmt.Println("person height is", p.Height)

	fmt.Println("app complete ...")

}
