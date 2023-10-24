package main

import (
	"fmt"
	"os"
)

func main() {
	//只读方法读取目录下的readme.txt
	file, err := os.Open("../readme.txt")
	if err != nil {
		fmt.Println("出错原因:", err)
		return
	}

	defer file.Close() //延迟关闭
	//读取文件
	var temp [128]byte
	v, err := file.Read(temp[:]) //基于数组得到切片
	if err != nil {
		fmt.Println("出错", err)
		return
	}
	fmt.Printf("读取了%d个字节\n", v)
	fmt.Println(string(temp[:]))
}
