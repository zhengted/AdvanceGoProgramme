package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

/*
	切片底层结构
	type SliceHeader struct {
		Data uintptr
		Len int
		Cap int
	}
*/

func main() {
	fmt.Println("-- slice defination --")
	// 切片的定义方式
	var (
		_ []int     // nil切片，和nil相等，一般用来表示一个不存在的切片
		_ = []int{} // 空切片，切片中没有元素，表示空集合

		c = []int{1, 2, 3} // len = 3 cap = 3
		_ = c[:2]          // len = 2 cap = 3
		_ = c[0:2:cap(c)]  // len = 2 cap = 3
		_ = c[:0]          // len = 0 cap = 3

		_ = make([]int, 2, 4) // len = 2 cap = 4
	)

	fmt.Println("-- slice add element --")
	// 切片追加元素的方式
	// 尾部添加
	var a []int
	a = append(a, 1)
	a = append(a, 1, 2, 3)
	a = append(a, []int{1, 2, 3}...) // 追加切片需要解包

	// 头部添加
	var b []int
	b = append([]int{0}, b...) // 头部添加元素需要把单个元素变成切片
	b = append([]int{-3, -2, -1}, b...)
	//在开头一般都会导致内存的重新分配，而且会导致已有的元素全部复制1次。
	//因此，从切片的开头添加元素的性能一般要比从尾部追加元素的性能差很多。

	// 中间添加
	a = append(a[:2], append([]int{3}, a[2:]...)...)       // 链式使用append 添加单个元素
	a = append(a[:2], append([]int{1, 2, 3}, a[2:]...)...) // 添加多个元素
	// 第二个append多了一次复制操作，增加了临时切片，影响效率。以下为优化代码

	// 添加单个元素
	a = append(a, 0)     // 切片扩展一个空间
	copy(a[2+1:], a[2:]) // 将2后的部分 复制到3
	a[2] = 123
	// 添加多个元素
	x := []int{4, 5, 6}
	a = append(a, x...)
	copy(a[2+len(x):], a[2:])
	copy(a[2:], x)

	fmt.Println("-- slice delete element -- ")
	N := 2

	// 移动指针删除数组元素
	// 删除尾部
	a = []int{1, 2, 3}
	a = a[:len(a)-1] // 删除尾部1个元素
	a = a[:len(a)-N] // 删除尾部N个元素
	// 删除头部
	a = []int{1, 2, 3}
	a = a[1:] // 删除开头1个元素
	a = a[N:] // 删除开头N个元素

	// 移动数据删除元素
	a = []int{1, 2, 3}
	a = append(a[:0], a[1:]...) // 删除开头1个元素
	a = append(a[:0], a[N:]...) // 删除开头N个元素
	a = []int{1, 2, 3}
	a = a[:copy(a, a[1:])] // 删除开头1个元素
	a = a[:copy(a, a[N:])] // 删除开头N个元素

}

// 避免切片的内存泄漏
func FindPhoneNumber(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return regexp.MustCompile("[0-9]+").Find(b)
}

//这段代码返回的[]byte指向保存整个文件的数组。
//因为切片引用了整个原始数组，导致自动垃圾回收器不能及时释放底层数组的空间。
//一个小的需求可能导致需要长时间保存整个文件数据。
//这虽然这并不是传统意义上的内存泄漏，但是可能会拖慢系统的整体性能。

// 正解：
func FindPhoneNumberEx(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	return append([]byte{}, b...)
}

// 同理在执行切片删除元素时也会发生内存泄漏
// 如果切片存放的是指针，删除元素时，底层数组依然在使用该元素。导致这个元素一直无法被回收
