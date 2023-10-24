package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 并发安全的map
var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	//对m执行20个并发的读写操作
	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)         //存储k-v
			value, _ := m.Load(key) //根据key读取值
			fmt.Printf("k=%v,value=%v", key, value)
			wg.Done()
		}(i) //匿名函数的调用，末尾加上(i)闭包，参数(n int)接受一个整数参数n,该参数是循环中的迭代变量i的值
	}
	wg.Wait()
}
