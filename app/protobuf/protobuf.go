package main

import (
	"fmt"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"mikeian-for-golang/app/protobuf/city"
	"mikeian-for-golang/app/protobuf/person"
)

func main() {

	fmt.Println("app start ... ")
	pa, _ := anypb.New(&person.Person{})
	fmt.Println("pa value is ", pa)
	c := city.City{
		Name:    "shenzhen",
		Address: "guangdongshen nanshanqu",
	}
	anyCity, _ := anypb.New(&c)
	fmt.Println("wrap any city value is ", anyCity)
	fmt.Println("concrete c:", c)
	uwc := &city.City{}
	anyCity.UnmarshalTo(uwc)
	fmt.Println("unwrap any city value2 is ", uwc)

	p := person.Person{
		Name:   "mikeian",
		Age:    20,
		City:   &c,
		Dv:     wrapperspb.Double(20),
		Any:    anyCity,
		Action: &person.Person_Start{Start: false},
	}

	fmt.Println(p)

	fmt.Println("app complete ... ")

}
