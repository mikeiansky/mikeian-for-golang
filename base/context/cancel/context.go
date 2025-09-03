package main

import (
	"context"
	"fmt"
)

func UseContext(ctx context.Context) {

	cd := <-ctx.Done()
	fmt.Println("context done", cd)

}

func main() {

	fmt.Println("app start ...")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		fmt.Println("cancel context")
		//time.Sleep(1 * time.Second)
		cancel()
	}()

	UseContext(ctx)

	fmt.Println("app complete ...")
}
