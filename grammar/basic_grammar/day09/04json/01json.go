package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Age    int
	gender string
}

func main() {
	lily := Student{
		Id:     1,
		Name:   "lily",
		Age:    20,
		gender: "女生",
	}
	//编码结构体变为字符串
	encodeInfo, err := json.Marshal(&lily)
	if err != nil {
		fmt.Println("json.Marshal,err :", err)
		return
	}
	fmt.Println("encodeInfo", string(encodeInfo))

	//解码反序列化
	var lily2 Student
	//将encodeInfo反序列化到lily2
	if err := json.Unmarshal([]byte(encodeInfo), &lily2); err != nil {
		fmt.Println("json.unmashal,err:", err)
		return
	}
	fmt.Println("name:", lily2.Name)
	fmt.Println("gender:", lily2.gender)
	fmt.Println("age:", lily2.Age)
	fmt.Println("id:", lily2.Id)

}
