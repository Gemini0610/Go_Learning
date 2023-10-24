package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	v := reflect.TypeOf(x) //拿到x的动态类型信息
	fmt.Printf("type:%v\n kind:%v\n", v.Name(), v.Kind())
}
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	K := v.Kind()
	switch K {
	case reflect.Int64:
		//v.Int()从反射中获取整型的原始值，然后通过int64强行类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		//v.Float()从反射中获取浮点的原始值，然后通过float32强行类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		//v.Float()从反射中获取浮点的原始值，然后通过float32强行类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}
func main() {
	var a *float32 //指针
	var b myInt    //自定义类型
	var c rune     //类型别名
	reflectType(a)
	reflectType(b)
	reflectType(c)

	type person struct {
		name string
		age  int
	}

	type book struct{ title string }
	var d = person{
		name: "小王子",
		age:  18,
	}
	var e = book{title: "香蕉斤斤计较"}
	reflectType(d)
	reflectType(e)

	var z float32 = 3.14
	var g int64 = 100
	reflectValue(z)
	reflectValue(g)
	n := reflect.ValueOf(10)
	fmt.Printf("type c :%T\n", n)
}
