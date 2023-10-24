package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { //input.scan读入下一行
		counts[input.Text()]++ //input.Text获取读取的内容
		// 等价于 line := input.Text()
		// counts[line] = counts[line]+1
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
