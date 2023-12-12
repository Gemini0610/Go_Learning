package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//1、建立连接
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("net.Dial err : ", err)
		return
	}
	fmt.Println("client与server建立连接")
	sendData := []byte("helloworld")
	for {
		//向服务器发送数据
		cnt, err := conn.Write(sendData)
		if err != nil {
			fmt.Println("conn.Write err :", err)
			return
		}
		fmt.Println("client ====> server cnt:", cnt, "data:", string(sendData))

		//接收服务器返回的数据
		//创建buf,用于接收服务器返回的数据
		buf := make([]byte, 1024)
		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}
		fmt.Println("Client <==== Server cnt :", cnt, "data:", string(buf[:cnt]))
		time.Sleep(1 * time.Second)
	}
	//关闭连接
	conn.Close()
}
