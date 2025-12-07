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
	region := city.City_Region{
		Id:   1,
		Name: "shenzhen",
	}
	// 先创建，再使用
	anyCity, _ := anypb.New(&c)
	fmt.Println("wrap any city value is ", anyCity)
	//fmt.Println("concrete c:", c)
	uwc := &city.City{}
	anyCity.UnmarshalTo(uwc)
	fmt.Println("unwrap any city value2 is ", uwc)

	children := []string{"ian", "Pop"}

	p := person.Person{
		Name:   "mikeian",
		Age:    20,
		City:   &c,
		Dv:     wrapperspb.Double(20),
		Any:    anyCity,
		Region: &region,
		Child:  children,
		//Action: &person.Person_Start{Start: false},
		//Action: &person.Person_Stop{Stop: true},
		Action: &person.Person_Update{Update: "tag-001"},
	}

	fmt.Println(p.Name)

	fmt.Println("app complete ... ")

}
