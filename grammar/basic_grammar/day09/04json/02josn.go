package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string `json:"-"`            //===>在使用json编码时，这个编码不参与
	Subject string `json:"Subject_name"` //==>编码后名称会变成这个
	Age     int    `json:"age,string"`   //==>将age的int类型变为string类型
}

func main() {
	t1 := Teacher{
		Name:    "DUKE",
		Subject: "golang",
		Age:     18,
	}

	fmt.Println("t1:", t1)
	encodeInfo, _ := json.Marshal(&t1)
	fmt.Println("encodeInfo:", string(encodeInfo))
}
