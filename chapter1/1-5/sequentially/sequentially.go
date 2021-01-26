package main

import (
	"time"
)

// 顺序一致性内存模型
func main() {
	done := make(chan bool)
	go func() {
		println("hello,world")
		done <- true // 这里用 done<-false和 close(done) 能起到一样的效果
	}()
	<-done
	mainEx()
}

// 多个goroutine的顺序控制
var limit = make(chan int, 3)
var work = []func(){
	func() { println("1"); time.Sleep(1 * time.Second) },
	func() { println("2"); time.Sleep(1 * time.Second) },
	func() { println("3"); time.Sleep(1 * time.Second) },
	func() { println("4"); time.Sleep(1 * time.Second) },
	func() { println("5"); time.Sleep(1 * time.Second) },
}
var doneEx chan bool

func mainEx() {
	for i, w := range work {
		i := i
		go func(tmpW func()) {
			limit <- 1
			tmpW()
			<-limit
			if i == len(work) {
				doneEx <- true
			}
		}(w) // 这里填入参数可以保证调用的tmpW和循环创建的w是同一个值，另一种方法是拷贝一份w
	}
	// 使用空select管道选择可以让main阻塞 同样可以用for <-done的方式让程序阻塞
	//select {
	//
	//}
	<-doneEx
}
