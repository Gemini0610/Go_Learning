package main

import (
	"fmt"
	"os"
)

// 在命令行中输入go run test2.go 123 输出 111123
func main() {
	s, sep := "", "111"
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
