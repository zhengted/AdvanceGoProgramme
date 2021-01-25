package main

import "fmt"

func main() {
	// 关于解包和不解包的接口参数
	var a = []interface{}{123, "abc"}
	Print(a...) // 123 abc	 解包	等价于 Print(123,abc)
	Print(a)    // [123,abc] 不解包	等价于 Print([]interface{}{123,abc})

	// 闭包访问外部变量
	// 错误写法：
	for i := 0; i < 3; i++ {
		defer func() { fmt.Println(i) }() //输出 333
	}
	// 正解
	for i := 0; i < 3; i++ {
		i := i // 定义一个循环体内局部变量i
		defer func() { println(i) }()
	}
	// 或者
	for i := 0; i < 3; i++ {
		// 通过函数传入i
		// defer 语句会马上对调用参数求值
		defer func(i int) { println(i) }(i)
	}
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}

type File struct {
	fd int
}

func (f *File) Read(offset int64, data []byte) int {
	return 0
}

func (f *File) Close() error {
	return nil
}

func TmpFileTest() {
	// 1.方法表达式特性 可以将对象方法还原为普通类型的函数
	var CloseFile = (*File).Close
	var ReadFile = (*File).Read
	ReadFile(&File{64}, 0, []byte{})
	CloseFile(&File{64})

	// 2.绑定对象
	f := &File{64}
	var Close = func() error {
		return (*File).Close(f)
	}
	var Read = func(offset int64, data []byte) int {
		return (*File).Read(f, offset, data)
	}
	Read(10, []byte{})
	Close()
}
