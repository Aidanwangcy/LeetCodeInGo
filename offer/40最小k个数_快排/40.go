package main

import (
	"fmt"
)

func main()  {
	arr := []int{4,5,1,6,2,7,3,8}
	k := 4
	ans := getLeastNumbers(arr, k)
	for _, v := range ans {
		fmt.Printf("%v ",v)
	}
	fmt.Print("\n")
}

func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 {
		return arr
	}
	low := 0
	high := len(arr) - 1
	//寻找基准值指针
	index := partition(arr, low, high)
	//如果基准值指针不等于k，继续寻找
	for index != k-1 {
		//基准值指针在k右边，说明index左边都比index所指值小，把high指针调整到index-1
		if index > k-1 {
			high = index - 1
			index = partition(arr, low, high)
		} else {
			//否则的话，寻找右半部分的基准值
			low = index + 1
			index = partition(arr, low, high)
		}
		//直到找到以k指向为基准值的元素，k左侧全部小于k
	}
	return arr[:k-1]
}

func partition(arr []int, low, high int) int {
	if low == high {
		return low
	}
	if low < high {
		//基准值
		base := arr[low]
		l := low
		r := high
		for {
			//如果最右>基准，指针左移，元素不动
			for l < r && arr[r] >= base {
				r--
			}
			//如果最左<基准，指针右移，元素不动
			for l < r && arr[l] <= base {
				l ++
			}
			//如果跨越就退出循环
			if l >= r {
				break
			}
			//找到了不符合条件的两个数，arr[l]大于基准值，arr[r]小于基准值，交换二者位置
			arr[l], arr[r] = arr[r], arr[l]
		}
		//循环结束后，将基准值放到中间位置
		arr[low], arr[l] = arr[l], arr[low]
		return l
	}
	return -1
}