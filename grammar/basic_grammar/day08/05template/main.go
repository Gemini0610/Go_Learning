package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// template示例
type User struct {
	UserName string
	Password string
	Age      int
}

func info(w http.ResponseWriter, r *http.Request) {
	//调用template
	//打开一个模板文件

	t, err := template.ParseFiles("./info.html")
	if err != nil {
		fmt.Println("parse html failed, err :", err)
		return
	}
	//用数据渲染模板
	// data := "<li>hhhh</li>"
	// t.Execute(w, data)

	// num := rand.Intn(10)
	// dataStr := string(data) //转换成字符串
	// if num > 5 {
	// 	dataStr = strings.Replace(dataStr, "{oo xx}", "<li>《SHW》</li>", 1)
	// } else {
	// 	dataStr = strings.Replace(dataStr, "{oo xx}", "<li>《哈利波特》</li>", 1)
	// }
	// w.Write([]byte(dataStr))

	//结构体
	user := User{
		"TOM",
		"123",
		12,
	}
	t.Execute(w, user)
}
func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe("127.0.0.1:8080", nil)
}
