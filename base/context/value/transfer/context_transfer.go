package main

import (
	"context"
	"fmt"
)

func UseContext(ctx context.Context) {
	value := ctx.Value("tag")
	fmt.Println("use context value is ", value)
	// update value
	context.WithValue(ctx, "tag", "use context value")
}

func main() {
	fmt.Println("context transfer start ... ")

	ctx := context.WithValue(context.Background(), "tag", "hello world")

	UseContext(ctx)

	tag := ctx.Value("tag")
	fmt.Println("tag at main func is ", tag)

	fmt.Println("context transfer complete ... ")

}
