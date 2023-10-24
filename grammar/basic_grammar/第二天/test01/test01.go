package main

import "fmt"

const boilingF = 212.0 //包一级范围声明语句

func main() {
	//在函数内部声明
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}
