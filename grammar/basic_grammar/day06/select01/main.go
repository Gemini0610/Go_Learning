package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	//打印出10以内的奇数
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
		//第二次for循环，i=2，由于通道缓冲区已经满了，初始容量设定为1
		//所以ch <- i这个分支不满足，而x := <- ch
		//这个分支可以执行从通道接收值1并赋值给变量x，在终端打印出1
	}
}
