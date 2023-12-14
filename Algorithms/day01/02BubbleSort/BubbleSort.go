package main

import "log"

func main() {
	arr := []int{10, 3, 0, -9, 2, 100, -22, 54, 23, 7}
	log.Println(arr)
	bubbleSort(arr)
	log.Println(arr)
}

func bubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}
	//因为冒泡排序是从0开始 01 12 这样比较大小，大的到最后固定了所以是end为for循环次数
	//首先定义n为数组的长度
	n := len(arr)
	//之后第一个for循环,是用来进行遍历大的轮廓的,由于数组从0开始计算所以end最大为n-1也就是初始值为n-1
	for end := n - 1; end > 0; end-- {
		//第二个for循环是用来找i与i+1的比较大小，并且i的最大值不能为end因为这样就不能进行比较了
		for i := 0; i < end; i++ {
			//若i的值大于i+1的值则进行互换位置
			if arr[i] > arr[i+1] {
				swap(arr, i, i+1)
			}
		}
	}
}

//互换值的函数
func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
