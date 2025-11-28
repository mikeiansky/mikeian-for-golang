package main

import "fmt"

type Tag01 struct {
	Name  string
	Count int32
}

type cf func(t *Tag01)

func executeConfig(t *Tag01, cfs ...cf) {
	for _, cf := range cfs {
		cf(t)
	}
}

// 不推荐这种方式，函数本身是有引用的语义的
func executeSimple(t *Tag01, cf *cf) {
	(*cf)(t)
}

func main() {

	fmt.Println("test type func start ... ")

	t := &Tag01{
		Name:  "test",
		Count: 1,
	}

	var f1 cf = func(t *Tag01) {
		fmt.Println("update 1 start")
		t.Count++
		t.Name = t.Name + "-1"
		fmt.Println("update 1 complete")
	}

	f2 := func(t *Tag01) {
		fmt.Println("update 2 start")
		t.Count++
		t.Name = t.Name + "-2"
		fmt.Println("update 2 complete")
	}

	ts := []cf{f1, f2}

	fmt.Println("before execute config ... ")
	fmt.Println(t)
	executeConfig(t, ts...)
	fmt.Println("after execute config ... ")
	fmt.Println(t)
	executeSimple(t, &f1)
	fmt.Println(t)
	executeSimple(t, (*cf)(&f2))
	fmt.Println(t)
	fmt.Println("test type func complete ... ")

}
