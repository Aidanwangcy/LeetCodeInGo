package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	ans := constructArr(a)
	for _, v := range ans {
		fmt.Printf("%v ", v)
	}
}

func constructArr(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	left, right := make([]int, len(a)), make([]int, len(a))
	left[0] = 1
	right[len(a)-1] = 1
	//左边到i的数的数组
	for i := 1; i < len(a); i++ {
		left[i] = left[i-1] * a[i-1]
	}
	//右边到i的数的数组
	for i := len(a) - 2; i >= 0; i-- {
		right[i] = right[i+1] * a[i+1]
	}
	//左右相乘，即为目标数组的值
	arr := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		arr[i] = left[i] * right[i]
	}
	return arr
}
