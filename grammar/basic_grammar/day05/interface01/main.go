package main

import "fmt"

//接口值
func main() {
	var x interface{} //接口存儲的值<type,value>
	var a int64 = 100
	var b int32 = 10
	var c int8 = 1
	x = a //<int64,100>
	x = b
	x = c
	fmt.Println(x)
	value, ok := x.(int8)
	if ok {
		fmt.Printf("x存的是一个int8类型,值为%v\n", value)
	} else {
		fmt.Printf("存的不是int8")
	}

	//类型断言
	//若猜对了，ok=true, value =对应类型的值
	//若猜错了，ok=false，value=对应类型的0值
	value1, ok := x.(float64)
	fmt.Printf("ok:%t value:%v, value type:%T\n", ok, value1, value1)
}
