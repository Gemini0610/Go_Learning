package main

import (
	"fmt"
)

func f1() int { //1.return=x 2.x++ 3.ret 结束===>5，因为没有改变return的值，只是改变了x的值,没有规定返回的变量名为x
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f2() (x int) { //1.return=x 2.x++ 3.ret===>6  前提规定好返回的变量叫x
	defer func() {
		x++
	}()
	return 5
}
func f3() (y int) { //1.return=y(5) 2.x++ 3.ret==>5
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) { //1.retrun=x 值为5 2.x++ defer内部的x进行x++ 3.ret==>5
		x++
	}(x)
	return 5
}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
