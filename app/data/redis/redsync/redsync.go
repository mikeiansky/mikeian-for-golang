package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

func doWork(ctx context.Context, mutex *redsync.Mutex, key, value, tag string, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println(tag, "start useKey:", key, value)
	time.Sleep(5 * time.Second)
	fmt.Println(tag, "complete useKey:", key, value)
	mutex.Unlock()
	wg.Done()
}

func main() {
	// Create a pool with go-redis (or redigo) which is the pool redsync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "192.168.31.109:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	// Create an instance of redsync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	mutexname := "my-global-mutex"
	mutex := rs.NewMutex(mutexname)

	//// Obtain a lock for our given mutex. After this is successful, no one else
	//// can obtain the same lock (the same mutex name) until we unlock it.
	//if err := mutex.Lock(); err != nil {
	//	panic(err)
	//}
	//fmt.Println("do my work start ... ")
	//
	//time.Sleep(5 * time.Second)
	//// Do your work that requires the lock.
	//fmt.Println("do my work complete ... ")
	//
	//// Release the lock so other processes or threads can obtain a lock.
	//if ok, err := mutex.Unlock(); !ok || err != nil {
	//	panic("unlock failed")
	//}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	ctx := context.Background()
	go doWork(ctx, mutex, mutexname, "first", "001", wg)
	go doWork(ctx, mutex, mutexname, "second", "002", wg)
	wg.Wait()
	fmt.Println("app complete")

}
