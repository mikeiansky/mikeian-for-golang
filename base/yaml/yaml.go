package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type App struct {
	Title   string `yaml:"title"`
	Port    int    `yaml:"port"`
	Company struct {
		Name  string `yaml:"name"`
		Count int    `yaml:"count"`
	}
}

func main() {
	fmt.Println("app start ... ")

	fn := "base/yaml/app.yaml"

	data, err := os.ReadFile(fn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(data))
	}

	app := App{}

	if err := yaml.Unmarshal(data, &app); err != nil {
		// before
		fmt.Println("[this word is analyse] yaml config error:", err)
		// after
	}

	fmt.Println("read app data:", app)
	fmt.Println("read app name:", app.Title)
	fmt.Println("read app port:", app.Port)
	fmt.Println("read app company name:", app.Company.Name)
	fmt.Println("read app company count:", app.Company.Count)

	fmt.Println("app complete ... ")
}
