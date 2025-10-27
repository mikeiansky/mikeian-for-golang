package main

import (
	"context"
	"fmt"
)

func main() {
	// 创建一个带有键值对的 context
	ctx := context.WithValue(context.Background(), "userID", 123)
	ctx = context.WithValue(ctx, "role", "admin")

	// 模拟一个函数使用这些值
	printUserInfo(ctx)
}

func printUserInfo(ctx context.Context) {
	userID := ctx.Value("userID")
	role := ctx.Value("role")

	fmt.Printf("用户ID: %v, 角色: %v\n", userID, role)
}
