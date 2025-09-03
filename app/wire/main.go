package main

import "fmt"

type App struct {
	Wc *WireConfig
}

func NewApp(wc *WireConfig) *App {
	return &App{
		Wc: wc,
	}
}

func main() {
	fmt.Println("app start ... ")
	app := createApp()
	fmt.Println("create app is", app)
	fmt.Println("app complete ...")
}
