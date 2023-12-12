package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//http包
	client := http.Client{}

	//func (c *Client) Get(url string) (resp *Response, err error)
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("client.Get, err:", err)
		return
	}

	body := resp.Body
	fmt.Println("body 111:", body)
	readBodyStr, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("body string : ", string(readBodyStr))
	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")

	fmt.Println("content-type:", ct)
	fmt.Println("Date:", date)
	fmt.Println("server:", server)

	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status

	fmt.Println("url:", url)
	fmt.Println("code:", code)
	fmt.Println("status:", status)

}
