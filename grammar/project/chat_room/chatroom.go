package main

import (
	"crypto/rand"
	"fmt"
	"net"
	"strings"
)

// 基础阶段,所以没有科学地写多个文件
type User struct {
	//名字
	name string
	//唯一的id
	id string
	//管道
	msg chan string
}

// 创建全局的map结构，用于保存所有用户 var 用于声明变量
var allUsers = make(map[string]User)

// 定义一个message全局通道，用于接收任何人发送过来消息
var message = make(chan string, 10)


func main() {
	//创建服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("连接出错, err:", err)
		return
	}
	//启动全局唯一的go程,负责监听message通道,写给所有用户
	go broadcast()

	fmt.Println("服务器启动成功，监听中.....")

	for {
		fmt.Println("主go程监听中...")
		//监听
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("监听错误,err:", err)
			return
		}

		//建立连接
		fmt.Println("建立连接成功！")
		//启动处理业务的goroutine
		go handler(conn)
	}
}

// 处理业务
func handler(conn net.Conn) {

	fmt.Println("启动业务")
	//客户端与服务器建立连接的时候,会有ip和port====>这里将其作为用户di
	clientAddr := conn.RemoteAddr().String()
	//创建user
	newUser := User{
		id:   clientAddr,            //id，作为在map中的key
		name: clientAddr,            //可以修改，会提供rename方法，建立连接，这个是初始值
		msg:  make(chan string, 10), //需要make空间，否则无法写入数据
	}

	//添加user到map结构
	allUsers[newUser.id] = newUser
	//定义一个退出信号，用于监听client退出
	var isQuit = make(chan bool)

	//启动go程，负责监听退出信号
	go watch(&newUser, conn, isQuit)
	//启动go程，将msg信息返回给客户端
	go writeBackToClient(&newUser, conn)

	//向message写入当前用户上线，这样就能够广播所有人
	loginInfo := fmt.Sprintf("[%s]:[%s]====>上线了", newUser.id, newUser.name)
	message <- loginInfo

	for {
		//具体业务逻辑

		buf := make([]byte, 1024)

		//读取客户端发来的请求数据
		cnt, err := conn.Read(buf)
		if cnt == 0 {
			fmt.Println("客户端退出")
			//map删除 用户 conn close
			//服务器还可以主动退出
			//这里不进行真正的退出，而是发送一个退出信号，统一做出退出处理
			//可以使用新的管道来实现信号传递
			isQuit <- true
		}
		if err != nil {
			fmt.Println("读取错误,err:", err)
			return
		}
		fmt.Println("服务器接收客户端发来的数据为：", string(buf[:cnt-1]), "cnt", cnt)

		//-------------业务逻辑处理 开始----------
		//1、查询当前所有用户 who
		//		首先，判断接收的数据是不是who  ===> 长度&&字符串

		userInput := string([:cnt-1]) //这是用户输入的数据，最后一个是回车，所以-1
		if len(userInput) == 4 && userInput == "\\who" {
			//		之后，遍历allUsers这个map， key := userid， value:= user本身，将id和name拼接成字符串返回给客户端

			fmt.Println("用户即将查询所有用户信息")
			//这个切片包含所有的用户信息
			var userInfos []string 
			for _, user := range allusers {
				userInfo := fmt.Sprintf("userid:%s, username:%s",user.id,user.name)
				userInfos = append(userInfos, userInfo)
			}
			//最终写到管道中一定是一个字符串
			r := strings.Join(userInfos, "\n")
		}else if len(userInput) > 9 && userInput[:7] == "\\rename"{
			//重命名
			//读取数据判断长度7，判断字符是rename
			//使用|进行分割，获取|后面的部分，作为名字
			newUser.name:= strings.Split(userInput, "|")[1]//前面是字符串，后面是使用什么来切分
			//更新用户名字newUser.name = Duke
			//通知客户端，更新成功
			newUser.msg <- "sucessfully"
		}else{
			//若用户输入的不是命令，只是普通的聊天信息，那么只需要写到广播通道中即可
			message <- userInput
			//将数据返回给查询的客户端
			newUser.msg <- r
		}

		//-------------业务逻辑处理 结束----------

	}
}

// 创建全局唯一的go程，用于向所有用户广播消息
func broadcast() {

	fmt.Println("广播go程启动成功....")
	defer fmt.Println("broadcast退出了.....")
	for {
		//1、从message中读取数据
		info := <-message
		//2、将数据写入到每一个用户的msg管道中
		for _, user := range allUsers {
			user.msg <- info
		}
	}
}

// 每个用户应该还有一个用来监听自己msg管道的go程，负责将数据返回给客户端
func writeBackToClient(user *User, conn net.Conn) {
	fmt.Printf("user:%s的go程在监听msg管道\n", user.name)
	for data := range user.msg {
		fmt.Printf("user:%s 写回给客户端的数据为%s\n", user.name, data)

		//Write(b []byte) (n int, err error)
		_, _ = conn.Write([]byte(data))
	}
}

//启动go程负责监听退出信号，触发后进行清理工作 delete map,close conn
func watch(user *User, conn net.Conn, isQuit <- chan bool)  {
	for{
		select {
		case <- isQuit:
			logoutInfo := fmt.Println("用户%s退出", user.name)
			delete(allUsers, user.id)
			//通知所有人
			message <- logoutInfo
			conn.Close()
			return
		
		case <- time.Ater(30*time.second):
			logoutInfo := fmt.Println("超时了，用户%s退出", user.name)
			delete(allUsers, user.id)
			//通知所有人
			message <- logoutInfo
			conn.Close()
			return
		}

	}
}
