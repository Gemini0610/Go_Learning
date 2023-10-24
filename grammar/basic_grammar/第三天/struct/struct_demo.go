package main

import "fmt"

//结构体
//创建一个新的类型，要用type关键字
type Student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var Sam = Student{
		name:   "Sam",
		age:    12,
		gender: "男",
		hobby:  []string{"足球", "篮球"},
	}
	fmt.Println(Sam)
	fmt.Println(Sam.name)

	var yami = new(Student)
	yami.name = "yami"
	yami.age = 12
	yami.gender = "女"
	fmt.Println(yami)
}
