package main

import (
	"encoding/json"
	"fmt"
)

// json序列化，元数据，意思是在json中代表什么
type student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var stu1 = student{
		ID:   12,
		Name: "jack",
		Age:  13,
	}
	fmt.Println(stu1)
	//序列化:json.Marshal将数据序列化为json格式的字符串
	//反序列化:json.Unmarshal
	v, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("false")
	}
	fmt.Println(string(v))
}
