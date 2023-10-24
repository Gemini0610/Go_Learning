package main

import "fmt"

type xiyiji interface {
	wash()
	shuaigan()
}
type Haier struct {
	name  string
	price float32
}

func (h Haier) wash() {
	fmt.Println("洗干净")
}
func (h Haier) shuaigan() {
	fmt.Println("晒干了")
}
func main() {
	var a xiyiji
	h1 := Haier{
		name:  "小神童",
		price: 32.3,
	}

	fmt.Printf("%T\n", h1)
	a = h1 //h1 实现了结构体haier的实例，haier这个结构体拥有两个方法，并且因为都是接口中的全部方法所以可以将h1赋值给a，并且输出,若仅仅实现一个方法那么a输出不了
	fmt.Println(a)
}
