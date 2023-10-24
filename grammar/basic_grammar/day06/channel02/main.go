package main

import "fmt"

//对于无缓冲的通道，要有一个接收方才能发送成功
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	//无缓冲通道ch的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作
	//这时10才能发送成功，两个goroutine将继续执行
	ch := make(chan int)
	go recv(ch) //创建一个goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
