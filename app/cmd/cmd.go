package main

import "fmt"
import "github.com/alexflint/go-arg"

type BaseCmd struct {
	Title string `arg:"--title,-t" help:"create title"`
	Size  int    `arg:"--size,-s" help:"create size"`
	Open  bool   `arg:"--open,-o" help:"open page"`
}

type ListCmd struct {
	Sort string `arg:"--sort,-r" help:"sort by"`
}

type TopCmd struct {
	BaseCmd
	List *ListCmd `arg:"subcommand:list"`
}

func main() {
	fmt.Println("app start ... ")
	var cmd TopCmd
	arg.MustParse(&cmd)
	fmt.Println(cmd)
	fmt.Println(cmd.List)
	fmt.Println("app complete ... ")
}
