package main

import "fmt"

type address struct {
	province string
	city     string
}

//匿名结构体
type student struct {
	name string
	string
	int
	addr address //结构体嵌套
}

func main() {
	var s1 = student{
		name:   "sam",
		string: "男",
		int:    12,
		addr: address{
			province: "辽宁",
			city:     "大连",
		},
	}
	fmt.Println(s1)
}
