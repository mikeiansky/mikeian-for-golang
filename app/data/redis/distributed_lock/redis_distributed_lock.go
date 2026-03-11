package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

func newClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.31.109:6379", // Redis 服务器地址
		Password: "",                    // 没有密码，默认为空
		DB:       0,                     // 默认使用 DB 0
	})
	return rdb
}

func useKey(ctx context.Context, rdb *redis.Client, key, value, tag string, wg *sync.WaitGroup) {
	fmt.Println(tag, "start useKey:", key, value)
	ret := rdb.SetNX(ctx, key, value, time.Second*10)
	fmt.Println(tag, "complete useKey:", key, value, ret)
	wg.Done()
}

func main() {
	rdb := newClient()
	ctx := context.Background()
	key := "test-lock-001"
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go useKey(ctx, rdb, key, "first", "001", wg)
	go useKey(ctx, rdb, key, "second", "002", wg)
	wg.Wait()
	fmt.Println("app complete")
}
