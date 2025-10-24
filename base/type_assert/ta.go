package main

import "fmt"

type IPerson interface {
	SayHello()
}

func useNormalPerson(person IPerson) {
	person.SayHello()
}

// 这里不能这样用
//func usePointPerson(person *IPerson) {
//	person.SayHello()
//}

type Animal interface {
}

type Normal struct {
}

type Temp struct {
}

func (n Normal) SayHello() {
	fmt.Println("hello normal")
}

func TestTypeIsPerson(val interface{}) {
	tn, ok := val.(IPerson)
	fmt.Println("test type is IPerson", ok, tn)
}

func TestTypeIsTemp(val interface{}) {
	tn, ok := val.(Temp)
	fmt.Println("test type is temp", ok, tn)
}

func testTypeIsAnimal(val interface{}) {

	tn, ok := val.(Temp)
	fmt.Println("test type is animal", ok, tn)
}

func main() {

	n := Normal{}
	n.SayHello()

	var np IPerson = n
	fmt.Println(np)

	var pp IPerson = &n
	fmt.Println(pp)

	nn, ok := np.(Normal)
	fmt.Println(ok, nn)

	pnn, pok := np.(*Normal)
	fmt.Println(pok, pnn)

	pn, ok := pp.(*Normal)
	fmt.Println(ok, pn)

	npn, pnok := pp.(Normal)
	fmt.Println(pnok, npn)
	fmt.Println("test type is temp")
	TestTypeIsTemp(np)
	TestTypeIsTemp(pp)

	fmt.Println("test type is animal")
	testTypeIsAnimal(np)
	testTypeIsAnimal(pp)

	fmt.Println("test type is person")
	TestTypeIsPerson(np)
	TestTypeIsPerson(pp)

	useNormalPerson(np)
	useNormalPerson(pp)
}
