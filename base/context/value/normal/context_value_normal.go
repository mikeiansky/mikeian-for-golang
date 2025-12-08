package main

import (
	"context"
	"fmt"
)

type ContextBody struct {
	Id      int32
	Name    string
	Address string
}

type ContextKey string

func CreateKey(key string) ContextKey {
	return ContextKey(key)
}

func main() {

	cb := ContextBody{
		Id:      1,
		Name:    "hello",
		Address: "shenzhen",
	}

	fmt.Println("cb :", cb)

	//// 创建一个带有键值对的 context
	ctx := context.WithValue(context.Background(), CreateKey("userID"), 123)
	ctx = context.WithValue(ctx, CreateKey("role"), "admin")
	ctx = context.WithValue(ctx, CreateKey("body"), &cb)
	//
	//// 模拟一个函数使用这些值
	ctx = printUserInfo(ctx)
	//
	// 外部获取新增的数据，这里强制转型需要判空
	innerValue := ctx.Value(CreateKey("inner")).(string)
	fmt.Println("innerValue:", innerValue)
}

func printUserInfo(ctx context.Context) context.Context {
	userID := ctx.Value(CreateKey("userID"))
	role := ctx.Value(CreateKey("role"))
	cb := ctx.Value(CreateKey("body"))
	ctb, _ := cb.(*ContextBody)
	fmt.Println("body.Id: ", ctb.Id)
	fmt.Println("body.Name: ", ctb.Name)
	fmt.Println("body.Address: ", ctb.Address)
	fmt.Printf("用户ID: %v, 角色: %v, body: %v\n", userID, role, cb)

	// 内部新增一个
	ctx = context.WithValue(ctx, CreateKey("inner"), "inner value ... ")

	return ctx
}
