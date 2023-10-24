package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("../readme.txt")
	if err != nil {
		fmt.Println("出错:", err)
		return
	}
	defer file.Close()
	//bufio缓冲区读取
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //读到换行停止
		if err == io.EOF {
			fmt.Print(str)
			return
		}
		if err != nil { 
			fmt.Println("出错", err)
			return
		}
		fmt.Print(str)
	}
}
