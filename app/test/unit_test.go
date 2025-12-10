package test

import (
	"fmt"
	"mikeian-for-golang/app/lib"
	"testing"
)

func TestHello(t *testing.T) {
	fmt.Println("test hello start ... ")
	tag := "hello"
	nt := lib.CreateObj(tag)
	fmt.Println(nt)
	fmt.Println("test hello complete ... ")
}

func TestHello2(t *testing.T) {
	fmt.Println("test hello2 start ... ")
	tag := "hello2"
	nt := lib.CreateObj(tag)
	fmt.Println(nt)
	fmt.Println("test hello2 complete ... ")
}
