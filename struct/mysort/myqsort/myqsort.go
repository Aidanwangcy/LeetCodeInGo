package myqsort

func QsortInt(nums []int) []int {
	qsort(nums, 0, len(nums)-1)
	return nums
}

func qsort(nums []int, left, right int) {
	if left < right {
		partitionIndex := partation(nums, left, right)
		qsort(nums, left, partitionIndex-1)
		qsort(nums, partitionIndex+1, right)
	}
}

func partation(nums []int, left, right int) int {
	pivot := left
	beginIndex := pivot + 1
	for i := beginIndex; i <= right; i++ {
		if nums[i] < nums[pivot] {
			swap(nums, i, beginIndex)
			beginIndex++
		}
	}
	//beginIndex-1的位置比pivot小所以更换到最左没有影响，保证左侧都为最小
	swap(nums, pivot, beginIndex-1)
	return beginIndex - 1
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
