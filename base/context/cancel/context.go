package main

import (
	"context"
	"fmt"
	"time"
)

func UseContext(ctx context.Context) {
	fmt.Println("use context start ... ")
	cd := <-ctx.Done()
	fmt.Println("context done", cd)

}

func main() {

	fmt.Println("app start ...")

	cb := context.Background()

	fmt.Println("cb ", cb)

	ctx, cancel := context.WithCancel(cb)

	go func() {
		fmt.Println("cancel context start")
		time.Sleep(1 * time.Second)
		cancel()
	}()

	UseContext(ctx)

	fmt.Println("app complete ...")
}
