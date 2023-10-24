package main

import "fmt"

//多返回值 value, ok := <-ch
//value 从通道中取出的值，如果通道被关闭则返回对应类型的零值
//ok 通道ch关闭时返回false 否则返回true
func f2(ch chan int) {
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("通道关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)
	f2(ch)
}
