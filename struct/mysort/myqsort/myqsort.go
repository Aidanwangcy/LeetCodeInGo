package myqsort

func QsortInt(nums []int) []int {
	qsort(nums, 0, len(nums)-1)
	return nums
}

func qsort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := partation(nums, left, right)
	qsort(nums, left, p-1)
	qsort(nums, p+1, right)
}

func partation(nums []int, left, right int) int {
	i, j := left, right
	p := left
	temp := nums[p]
	for i < j {
		for i <= p && nums[i] <= temp {
			i++
		}
		if i <= p {
			nums[p] = nums[i]
			p = i
		}
		for j >= p && nums[j] >= temp {
			j--
		}
		if j >= p {
			nums[p] = nums[j]
			p = j
		}
	}
	nums[p] = temp
	return p
}


