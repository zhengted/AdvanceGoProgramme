package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total struct {
	sync.Mutex
	value int
}

var valueEx uint64

// 如何使用原子性
func worker(wg *sync.WaitGroup, index int) {
	defer wg.Done() // 当前的goroutine执行完毕调用wg.Done
	var i uint64
	for i = 0; i <= 100; i++ {
		total.Lock()
		total.value += 1
		total.Unlock()

		// 以上三句可以修改成
		atomic.AddUint64(&valueEx, i)

		fmt.Printf("index:%d\ttotal.value:%d\tvalueEx:%d\n", index, total.value, valueEx)
	}
}

func main() {
	// wg的使用方法
	var wg sync.WaitGroup
	wg.Add(2) // 告诉wg需要等待两个goroutine
	go worker(&wg, 1)
	go worker(&wg, 2)
	wg.Wait() // 阻塞等待对应的goroutine执行完毕

	fmt.Println(total.value)
}
