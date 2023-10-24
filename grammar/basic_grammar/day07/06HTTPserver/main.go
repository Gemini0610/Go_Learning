package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// http.ResponseWriter用于写入HTTP响应，*http.Request表示HTTP请求的信息
func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() //在函数执行完毕后自动执行
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}
func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}
func main() {
	//http.HandleFunc("/", getHandler)
	http.HandleFunc("/", postHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}

}
