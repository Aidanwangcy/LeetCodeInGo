package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	ans := spiralOrder(matrix)
	for _, v := range ans {
		fmt.Printf("%v ", v)
	}
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
	l1, l2 := len(matrix), len(matrix[0])
	var ans []int
	top, bottom, left, right := 0, l1-1, 0, l2-1
	for top < bottom && left < right {
		for i := left; i < right; i++ {
			ans = append(ans, matrix[top][i])
		}
		for i := top; i < bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		for i := right; i > left; i-- {
			ans = append(ans, matrix[bottom][i])
		}
		for i := bottom; i > top; i-- {
			ans = append(ans, matrix[i][left])
		}
		right--
		top++
		bottom--
		left++
	}

	if top == bottom {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
	} else if left == right {
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][left])
		}
	}

	return ans
}
