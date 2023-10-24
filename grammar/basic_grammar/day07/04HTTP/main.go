package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get实例
func main() {
	resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err) //获取失败报出错误
		return
	}
	defer resp.Body.Close()                //关闭连接
	body, err := ioutil.ReadAll(resp.Body) //从resp.Body中读取数据并将其存储到body变量中
	if err != nil {
		fmt.Printf("read from resp.Body failed, err%v\n", err)
		return
	}
	fmt.Println(string(body))
}
