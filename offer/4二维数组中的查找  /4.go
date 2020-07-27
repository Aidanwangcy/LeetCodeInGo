package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		l := len(matrix[i])
		if l > 0 && matrix[i][0] <= target && target <= matrix[i][l-1] {
			ok := BinarySearch(matrix[i], target)
			if ok {
				return true
			}
		}
	}
	return false
}

//BinarySearch 二叉树搜索
func BinarySearch(a []int, target int) bool {
	l := 0
	r := len(a) - 1
	for l <= r {
		mid := l + (r-l)>>1 //右移除2
		if a[mid] == target {
			return true
		} else if a[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
