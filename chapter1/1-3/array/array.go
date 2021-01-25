package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func main() {
	// defination
	fmt.Println("-- Golang array defination --")
	var (
		a [3]int                    // 定义长度为3的int型数组, 元素全部为0
		b = [...]int{1, 2, 3}       // 定义长度为3的int型数组, 元素为 1, 2, 3
		c = [...]int{2: 3, 1: 2}    // 定义长度为3的int型数组, 元素为 0, 2, 3
		d = [...]int{1, 2, 4: 5, 6} // 定义长度为6的int型数组, 元素为 1, 2, 0, 0, 5, 6
		// 使用...进行初始化都是根据元素的数目自动计算
		// c数组因为0没有定义  初始化为 0
	)
	fmt.Printf("a :%v\n", a)
	fmt.Printf("b :%v\n", b)
	fmt.Printf("c :%v\n", c)
	fmt.Printf("d :%v\n", d)

	// array and pointer
	fmt.Println("-- Array and pointer --")
	var e = [...]int{1, 2, 3}
	var f = &e
	fmt.Printf("e[0]:%d\te[1]:%d\n", e[0], e[1])
	fmt.Printf("f[0]:%d\tf[1]:%d\n", f[0], f[1])
	// f是e的指针 但是可以通过和e一样的访问方式访问数组
	// 注意：数组的长度是数组类型的组成部分，只想不同长度的数组指针类型也是完全不同

	// array traversal
	fmt.Println("-- Array traversal --")
	for i := range a {
		fmt.Printf("a[%d]:%d\n", i, e[i])
	}
	for i, v := range b {
		fmt.Printf("b[%d]:%d\n", i, v)
	}
	for i := 0; i < len(c); i++ {
		fmt.Printf("c[%d]:%d\n", i, c[i])
	}
	var times [5][0]int
	for range times {
		fmt.Println("hello,world") // 重复五次
	}
	// other type array
	fmt.Println("-- Other type array --")
	var (
		_ = [2]string{"1hello", "world"}
		_ = [...]string{"2hello", "world"}
		_ = [...]string{1: "3hello", 2: "world"}

		_ [2]image.Point
		_ = [...]image.Point{image.Point{0, 0}, image.Point{1, 1}}
		_ = [...]image.Point{{0, 0}, {1, 1}}

		_ [2]func(reader io.Reader) (image.Image, error)
		_ = [...]func(reader io.Reader) (image.Image, error){
			png.Decode,
			jpeg.Decode,
		}

		_ [2]interface{}
		_ = [...]interface{}{123, "hello"}

		_ = [2]chan int{}
	)

}
