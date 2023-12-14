package main

import "log"

func main() {
	arr := []int{10, 3, 8, -9, 2, 100, -22, 54, 23, 7}
	log.Println(arr)
	selectSort(arr)
	log.Println(arr)
}

func selectSort(arr []int) {
	//首先判断这个数组是否能够排序，为空不行，小于2也不行
	if arr == nil || len(arr) < 2 {
		return
	}
	//n是arr的长度
	n := len(arr)
	//之后先用第一个for循环进行遍历查找也就是笔记上的第一遍0-5，第二遍1-5以此类推
	//n-1的原因是在于5-5不需要再进行比较了
	for i := 0; i < n-1; i++ {
		//定义位置,第i个位置是该阶段最小值的位置
		minIndex := i
		//下一个for循环进行比较并将最小的位置定义为那里,这里j要设定为i的下一个位置这样才能进行比较,并且j要小于n
		for j := i + 1; j < n; j++ {
			//值若小，则将最小值的位置设定为j
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//进行交换的函数
		swap(arr, i, minIndex)
	}

}

//进行交换值的函数
func swap(arr []int, i int, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
