package main

import "fmt"

type animal interface {
	speak()
	move()
}
type Cat struct {
	name string
}

// cat类型要实现animal接口
func (c Cat) speak() {
	fmt.Println("miao")
}
func (c Cat) move() {
	fmt.Println("跑")
}

// 实现接口，使用指针接受者，和值接收者的区别
func main() {
	var x animal
	c1 := Cat{
		name: "小花",
	}
	fmt.Println(c1)
	x = c1
	x.move()
	fmt.Println(x)
	x.speak()

	tom := &Cat{ //tom是一个指针
		name: "tom",
	}
	(*tom).move()

}
