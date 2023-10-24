package main

import "fmt"

type Student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func newStudent(n string, age int, g string, h []string) Student {
	return Student{
		name:   n,
		age:    age,
		gender: g,
		hobby:  h,
	}
}
func main() {
	hobbySlice := []string{"篮球", "足球"}
	Sam := newStudent("sam", 12, "男", hobbySlice)
	fmt.Println(Sam)
}
