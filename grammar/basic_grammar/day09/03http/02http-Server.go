package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//注册路由
	//xxxx/user ====> func1
	//xxxx/abc  ====> func2
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		//request ===> 包含客户端发来的数据
		fmt.Println("用户请求详情:")
		fmt.Println("request:", r)
		//writer ===> 通过writer将数据返回给客户端
		_, _ = io.WriteString(w, "这是/user返回的数据")
	})
	http.HandleFunc("/abc", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "这是/abc返回的数据")
	})

	//func ListenAndServe(addr string, handler Handler) error
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		fmt.Println("http start failed,err:", err)
		return
	}
}
