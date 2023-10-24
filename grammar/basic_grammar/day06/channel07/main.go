package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello() {
	defer wg.Done()
	fmt.Println("hello world")
}
func main() {
	wg.Add(1)
	go hello() //启动另外一个goroutine执行hello函数
	fmt.Println("main hello world")
	wg.Wait()
}
