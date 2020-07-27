package mymergesort

func MergeSort(nums []int) []int {
	l := len(nums)
	if l<2 {
		return nums
	}
	mid := l/2
	left := nums[0:mid]
	right := nums[mid:]
	return merge(MergeSort(left), MergeSort(right))
}

func merge(left []int, right []int) []int {
	var ans []int
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			ans = append(ans, left[0])
			left = left[1:]
		} else {
			ans = append(ans, right[0])
			right = right[1:]
		}
	}
	for len(left) != 0 {
		ans = append(ans, left[0])
		left = left[1:]
	}
	for len(right) != 0 {
		ans = append(ans, right[0])
		right = right[1:]
	}
	return ans
}
