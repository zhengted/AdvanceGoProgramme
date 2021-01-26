package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 原子操作配合互斥锁性能代价太大，改为使用数字型标记
type singleton struct{}

var (
	instance    *singleton
	initialized uint32
	mu          sync.Mutex
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if instance == nil {
		defer atomic.StoreUint32(&initialized, 1)
		instance = &singleton{}
	}
	return instance
}

// 通用代码sync.Once的实现
var once sync.Once

func InstanceEx() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

// 简易的生产者消费者模型
// 后台线程始终读取最新的配置信息（生产者）
// 【消费者】始终使用最新的配置信息处理请求
func loadConfig() interface{} {
	return time.Now().UnixNano()
}

var req chan int

func requests() chan int {
	return req
}

func main() {
	var config atomic.Value // 保存当前配置信息

	// 初始化配置信息
	config.Store(loadConfig())

	// 启动一个后台线程, 加载更新后的配置信息
	go func() {
		for {
			time.Sleep(time.Second)
			config.Store(loadConfig())
			req <- 1
		}
	}()

	// 用于处理请求的工作者线程始终采用最新的配置信息
	for i := 0; i < 10; i++ {
		go func() {
			for r := range requests() {
				c := config.Load()
				fmt.Println(r, c)
			}
		}()
	}
	time.Sleep(time.Second * 10)
}
