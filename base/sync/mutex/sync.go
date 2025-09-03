package main

import (
	"fmt"
	"sync"
)

func NoSyncCount(size int) {
	sc := sync.WaitGroup{}
	sc.Add(2)
	count := 0

	go func() {
		for i := 0; i < size; i++ {
			count++
		}
		sc.Done()
	}()

	go func() {
		for i := 0; i < size; i++ {
			count++
		}
		sc.Done()
	}()

	sc.Wait()
	fmt.Println("no sync count", count)
}

func SyncCount(size int) {
	sm := sync.Mutex{}
	sc := sync.WaitGroup{}
	sc.Add(2)
	count := 0

	go func() {
		for i := 0; i < size; i++ {
			sm.Lock()
			count++
			sm.Unlock()
		}
		sc.Done()
	}()

	go func() {
		for i := 0; i < size; i++ {
			sm.Lock()
			count++
			sm.Unlock()
		}
		sc.Done()
	}()

	sc.Wait()
	fmt.Println("sync count", count)
}

func main() {
	fmt.Println("app start ...")
	size := 10000000
	NoSyncCount(size)
	SyncCount(size)
	fmt.Println("app complete ... count")
}
