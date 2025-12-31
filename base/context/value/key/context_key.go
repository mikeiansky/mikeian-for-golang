package main

import (
	"context"
	"fmt"
)

type authKey string

const (
	authInfoKey authKey = "authInfo"
)

type AuthInfo struct {
	MerchantId int64
}

func SaveAuthInfo(ctx context.Context, merchantId int64) context.Context {
	ctx = context.WithValue(ctx, authInfoKey, &AuthInfo{MerchantId: merchantId})
	return ctx
}

func GetAuthInfo(ctx context.Context) *AuthInfo {
	v := ctx.Value(authInfoKey)
	if v == nil {
		return nil
	}
	return v.(*AuthInfo)
}

func main() {
	ctx := SaveAuthInfo(context.Background(), 11223344)
	ev := GetAuthInfo(ctx)
	fmt.Println("exists value ", ev)

	nv := GetAuthInfo(context.Background())
	fmt.Println("nil value ", nv)

}
