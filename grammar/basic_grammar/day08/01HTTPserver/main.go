package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("hello, shw！"))
	//从hello.txt文件中读取数据到w中
	data, err := ioutil.ReadFile("./hello.html")
	if err != nil {
		fmt.Println("read from file failed, err:", err)
		return
	}
	w.Write(data)
}
func search(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hhhhhh!"))
}

// HTTP server端
func main() {
	http.HandleFunc("/hello", sayHello)
	http.HandleFunc("/web", search)
	//启动服务
	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		panic(err)
	}
}
