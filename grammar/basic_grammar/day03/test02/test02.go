package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "how do you do"
	//查看有哪些单词，依据" "进行分段
	wordSlice := strings.Split(s, " ")
	fmt.Println(wordSlice)
	wordMap := make(map[string]int, len(wordSlice))
	for _, word := range wordSlice {
		v, ok := wordMap[word] //ok 判断是否存在返回bool值
		fmt.Println(ok)
		if ok {
			wordMap[word] = v + 1
		} else {
			wordMap[word] = 1
		}
	}
	for k, v := range wordMap {
		fmt.Println(k, v)
	}
}
