package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func search(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("./hello.html")
	if err != nil {
		fmt.Println("read html failed, err:", err)
	}
	w.Write(data)
}
func index(w http.ResponseWriter, r *http.Request) {
	//r代表跟请求相关的所有内容
	//获取注册的信息
	//获取请求的方法
	fmt.Println("请求的方法", r.Method)
	r.ParseForm() //解析
	fmt.Printf("输出表单内容%v\n", r.Form)
	usnername := r.Form.Get("username")
	pwdValue := r.Form.Get("pwd")
	fmt.Println(usnername, pwdValue)
	w.Write([]byte("index"))
}

// HTTP server端
func main() {
	http.HandleFunc("/web", search)
	http.HandleFunc("/index", index)
	//启动服务
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}
