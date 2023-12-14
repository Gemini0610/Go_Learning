package main

import "log"

func main() {
	arr := []int{10, 3, 0, -9, 2, 100, -22, 54, 23, 7}
	log.Println(arr)
	InsertSort(arr)
	log.Println(arr)
}

func InsertSort(arr []int) {
	//判断是否可以进行排序
	if arr == nil || len(arr) < 2 {
		return
	}
	//不要忘记设定n为arr的长度
	n := len(arr)
	//第一个循环是负责设定范围的，从0-1开始所以i:=1
	for i := 1; i < n; i++ {
		//第二个循环是用来比较大小的  j是i的左边的位置,j+1与i是一样的意义,j--用于挨个比较大小
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			swap(arr, j, j+1)
		}
	}
}

func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
