package myheapsort

func HeapSort(nums []int) []int {
	numsLen := len(nums)
	buildMaxHeap(nums, numsLen)
	for i := numsLen-1; i >= 0; i-- {
		swap(nums, 0, i)
		numsLen--
		heapify(nums, 0, numsLen)
	}
	return nums
}

func buildMaxHeap(nums []int, numsLen int) {
	for i := numsLen/2 ; i >= 0; i-- {
		heapify(nums, i, numsLen)	
	}
}

func heapify(nums []int, upHeap, numsLen int)  {
	//left,right 指向最右下两个子节点
	left := 2*upHeap + 1
	right := 2*upHeap + 2
	largest := upHeap
	//判断堆尾和堆首谁大，堆尾大的话换到堆首
	if left < numsLen && nums[left] > nums[largest] {
		largest = left
	}

	if right < numsLen && nums[right] > nums[largest] {
		largest = right
	}
	if largest != upHeap {
		swap(nums, largest, upHeap)
		heapify(nums, largest, numsLen)
	}
}

func swap(nums []int, a, b int)  {
	nums[a], nums[b] = nums[b], nums[a]
}