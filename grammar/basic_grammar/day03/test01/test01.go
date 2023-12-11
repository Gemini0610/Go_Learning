package main

import "fmt"

func main() {

	b := [...]string{"beijing", "shanghai", "shenzhen"}
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	for index, value := range b { //根据index索引，取对应index的value值
		fmt.Println(index, value)
	}
	b[1] = "dalian"
	fmt.Println(b[1])

	//切片
	a1 := []int{1, 2, 3}
	c1 := a1
	a1[1] = 20
	fmt.Println(a1)
	fmt.Println(c1)

	s5 := make([]bool, 2, 10)
	fmt.Println(s5)
	//拷贝
	a2 := []int{1, 2, 3}
	b2 := make([]int, 3)
	copy(b2, a2)
	a2[2] = 100
	fmt.Println(a2)
	fmt.Println(b2)
}
