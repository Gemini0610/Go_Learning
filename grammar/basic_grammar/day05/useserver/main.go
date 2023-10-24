package main

import (
	"learninggolang/grammar/basic_grammar/day05/mylogger"
)

func main() {
	logger := mylogger.NewFileLogger("./", "XXX.log")
	stu1 := "Sam"
	logger.Debug("%s is a good man", stu1)

}
