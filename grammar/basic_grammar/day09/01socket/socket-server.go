package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//接收一个连接，只能执行一次
	//1、创建监听
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)
	listener, err := net.Listen("tcp", address)
	//net.Listen("tcp", "8848") 简写，冒号前面默认是本机127.0.0.1
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	fmt.Println("监听中")
	//2、Accept返回两个值 conn err
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("litener connect failed, err : ", err)
		return
	}
	fmt.Println("连接建立成功")
	//创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024) //使用make来创建字节切片,byte==>uint8
	//Read(b []byte) (n int, err error)
	//3、cnt真正读取client发来的数据长度
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("Client ====> Server,长度：", cnt, ", 数据：", string(buf))
	//4、服务器对客户端请求进行响应
	//将数据转成大写"hello">>>>"HELLO"
	//func ToUpper(s string) string
	upperData := strings.ToUpper(string(buf))

	//5、写数据
	//Write(b []byte) (n int, err error)
	cnt, err = conn.Write([]byte(upperData))
	fmt.Println("Client <==== Server,长度:", cnt, "数据:", upperData)
	//关闭连接
	conn.Close()
}
