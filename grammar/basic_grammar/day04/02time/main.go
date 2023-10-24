package main

import (
	"fmt"
	"time"
)

// time包练习题
// 编写一个函数，接收Time类型函数，函数内部由"2023/07/19 17:30:3"格式
func printTime(t time.Time) {
	timeStr := t.Format("2006/01/02 15:04:05")
	fmt.Println(timeStr)
}
func calTime() {
	start := time.Now().UnixNano() / 1000 //获取当前纳秒级时间戳，/1000是为了变成微妙
	fmt.Println("jjjj")
	time.Sleep(time.Millisecond * 30)
	end := time.Now().UnixNano() / 1000
	fmt.Printf("耗费了%d微妙\n", end-start)

}
func main() {
	now := time.Now()
	printTime(now)
	calTime()
}
