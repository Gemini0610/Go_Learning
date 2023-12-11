package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string //定义了s sep 两个字符串
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i] //s的旧值与 sep和os.Args[i]连接后赋值回
		sep = ""
	}
	fmt.Println(s)
}
