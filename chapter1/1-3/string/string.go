package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	// 字符串结构由两个信息组成：1是字符串指向的底层字节数组，2是字符串的字节的长度

	// 字符串的切片操作
	fmt.Println("-- String slice-type operation --")
	s := "hello, world"
	hello := s[:5]
	world := s[7:]

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]

	fmt.Println("hello:", hello)
	fmt.Println("world:", world)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)

	// 字符串底层
	fmt.Println("-- string basement --")
	fmt.Printf("%#v\n", []byte("Hello, 世界"))
	fmt.Printf("%v\n", []byte("Hello, 世界"))

	fmt.Printf("%#v\n", []rune("世界"))             // []int32{19990, 30028}
	fmt.Printf("%#v\n", string([]rune{'世', '界'})) // 世界
}

// 以下四种转换都因为底层结构存在差异，因此都会出现额外空间的开销
func str2bytes(s string) []byte { // 等同于[]byte(s)
	p := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		p[i] = c
	}
	return p
}

func bytes2str(s []byte) (p string) { // 等同于string(bytes)
	data := make([]byte, len(s))
	for i, c := range s {
		data[i] = c
	}
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&p))
	hdr.Data = uintptr(unsafe.Pointer(&data[0]))
	hdr.Len = len(s)

	return p
}

func str2runes(s string) []rune { // 等同于[]rune(str)
	var p []int32
	for len(s) > 0 {
		r, size := utf8.DecodeRuneInString(s)
		p = append(p, int32(r))
		s = s[size:]
	}
	return []rune(p)
}

func runes2string(s []int32) string { // 等同于string(runes)
	var p []byte
	buf := make([]byte, 3)
	for _, r := range s {
		n := utf8.EncodeRune(buf, r)
		p = append(p, buf[:n]...)
	}
	return string(p)
}
