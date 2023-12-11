package main

import (
	"fmt"
	"os"
)

// 0.定义结构体
type book struct {
	name    string
	writer  string
	price   float32
	publish bool
}

// 定义书籍存储切片
// type allBooks = make([]*book, 0, 200)
func showMenu() {
	fmt.Println("欢迎来到界面")
	fmt.Println("1.添加书籍")
	fmt.Println("2.修改书籍")
	fmt.Println("3.展示所有书籍")
	fmt.Println("4.退出")
}
func addBook() {
	var (
		name    string
		writer  string
		price   float32
		publish bool
	)
	fmt.Println("请根据信息输入")
	fmt.Println("请输入书名")
	fmt.Scanln(&name)
	fmt.Println("请输入作者")
	fmt.Scanln(&writer)
	fmt.Println("请输入价格")
	fmt.Scanln(&price)
	fmt.Println("请输入是否上架[ture|false]")
	fmt.Scanln(&publish)
}
func alterBook() {

}
func showBook() {

}

func main() {
	for {
		showMenu()
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			addBook()
		case 2:
			alterBook()
		case 3:
			showBook()
		case 4:
			os.Exit(0)
		}
	}

}
