package main

import (
	"fmt"
	"sync"
)

// 启动goroutine
// sync包中的waitgroup，sync并发原语
var wg sync.WaitGroup

func hello() {
	fmt.Println("hello world")
	wg.Done() //goroutine运行结束就登记-1
}
func main() {
	wg.Add(1)  //启动一个goroutine就登记+1
	go hello() //1.启动一个goroutine 2.在新的goroutine中执行hello函数
	fmt.Println("hello main")
	//time.Sleep(time.Second)
	//等hello执行完，执行hello的函数那个goroutine执行完。
	wg.Wait() //阻塞，一直等待所有的goroutine执行完毕
}
