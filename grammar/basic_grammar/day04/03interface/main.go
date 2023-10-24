package main

import "fmt"

type Cat struct{}

func (c Cat) Say() string {
	return "miao"
}

func (d Dog) Say() string { //(d Dog)函数的接受者
	return "wang"
}

type Dog struct{}
type Sayer interface {
	Say() string
}

func main() {
	var animalList []Sayer //接口是一个类型所以用切片表示
	c := Cat{}
	d := Dog{}
	animalList = append(animalList, c, d)
	fmt.Println(animalList)
}
