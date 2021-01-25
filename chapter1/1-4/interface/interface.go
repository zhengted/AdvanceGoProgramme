package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// 定制输出对象
type UpperWriter struct {
	io.Writer
}

func (p *UpperWriter) Write(data []byte) (n int, err error) {
	// 将原有的writer封装了一层方法，理解成中间件即可
	return p.Writer.Write(bytes.ToUpper(data))
}

func main() {
	fmt.Fprintln(&UpperWriter{os.Stdout}, "hello,world")
}
