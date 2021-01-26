package main

import (
	"fmt"
	"sync"
)

func doubleLock() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("hello,world")
		mu.Unlock()
	}()
	mu.Lock() // 这里执行的时候会被阻塞，要等到goroutine的unlock被执行才能执行这里
}
