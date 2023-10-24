package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://127.0.0.1:9090/post"
	contentType := "application/json"
	data := `{"name":"哈哈哈","age":20}`
	//url是目标服务器的url地址，post请求发送到这里
	//contentType是http请求的内容类型
	//strings.NewReader(data)用于创建一个包含请求主体数据的io.Reader接口，data是post请求主体发送的数据
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
