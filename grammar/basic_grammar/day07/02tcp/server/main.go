package main

import (
	"bufio"
	"fmt"
	"net"
)

// 一个TCP服务端，可以同时连接很多个客户端，Go中goroutine实现并发非常方便和高效，所以可以每建立一次链接，就创建一个goroutine去处理
// 把一个任务包装成一个函数来执行
// 单独处理链接的函数
func process(conn net.Conn) {
	defer conn.Close() //关闭链接
	for {
		reader := bufio.NewReader(conn)
		//数组
		var buff [128]byte
		n, err := reader.Read(buff[:]) //读取数据内容格式是数组，返回一个n和err
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}
		recvStr := string(buff[:n]) //将读到的内容格式化为string格式
		fmt.Println("收到client端发来的数据:", recvStr)
		conn.Write([]byte(recvStr)) //write方法调用将数据写入conn所代表的网络连接，[]byte(recvStr)将recvStr转换为byte数组，传递给write方法。
	}
}
func main() {
	//1、监听接口
	listen, err := net.Listen("tcp",
		"127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen:20000失败,err:", err)
		return
	}
	// 2、接收客户端请求建立连接
	//由于很多个端口用for循环
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("accpet failed, err:", err)
			continue //结束了
		}
		// 3、创建goroutine处理连接
		go process(conn) //启动一个goroutine处理链接，再写一个process处理函数
	}
}
