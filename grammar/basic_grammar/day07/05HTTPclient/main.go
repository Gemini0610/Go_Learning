package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	apiUrl := "http://127.0.0.1:9090/get"
	//URL param
	data := url.Values{} //用于构建HTTP请求的查询参数
	data.Set("name", "小王子")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl) //解析aipUrl，将其解析为一个URL对象u
	if err != nil {
		fmt.Printf("parse url requestUrl failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode() //对u中的查询参数进行url编码，将编码后的查询字符串设置u.RawQuery
	fmt.Println(u.String())
	resp, err := http.Get(u.String()) //向经过编码的URL发起HTTP GET请求，获取HTTP响应。
	if err != nil {
		fmt.Printf("post failed,err:%v\n", err)
		return
	}
	defer resp.Body.Close() //请求完成关闭响应体
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))

}
