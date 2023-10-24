package main

import (
	"fmt"
	"sync"
)

var (
	x  int64
	wg sync.WaitGroup //等待组
	m  sync.Mutex
)

func add() {
	//在for循环中进行加锁
	for i := 1; i <= 5000; i++ {
		//修改前加锁
		m.Lock()
		x = x + 1
		//修改后释放锁
		m.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(x)
}
