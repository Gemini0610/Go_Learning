package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// template示例
func info(w http.ResponseWriter, r *http.Request) {
	//打开一个模板文件
	t, err := template.ParseFiles("./info.html")
	if err != nil {
		fmt.Println("parse html file failed, err ", err)
		return
	}
	//用数据渲染模板
	data := "<li>hhhhhhhh</li>"
	t.Execute(w, data)
	// num := rand.Intn(10)
	// dataStr := string(data) //转换成字符串
	// if num > 5 {
	// 	dataStr = strings.Replace(dataStr, "{ooxx}", "<li>hhhhhhhh</li>", 1)
	// } else {
	// 	dataStr = strings.Replace(dataStr, "{ooxx}", "<li>wwwwwwww</li>", 1)
	// }
	// w.Write([]byte(dataStr))
}
func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
