package main


type MedianFinder struct {
	minArr []int	//小顶堆
	maxArr []int	//大顶堆
}

func Constructor() MedianFinder {
	heap := MedianFinder{
		minArr: make([]int, 0),
		maxArr: make([]int, 0),
	}
	return heap
}

func (this *MedianFinder) AddNum(num int) {
	//第一个元素，放到小顶堆
	if len(this.minArr) == 0 {
		this.minHeapUp(num)
		return
	}
	// 元素比小根堆的堆顶元素大
	if num > this.minArr[0] {
		// 判断如果将该元素扔到小根堆时，如果此时小根堆长度比大根堆大
		if len(this.minArr) > len(this.maxArr) {
			this.maxHeapUp(this.minArr[0])	// 将小根堆的堆顶元素扔到大根堆
			this.minHeapDown(num)			// 该元素放在小根堆的堆顶
		} else {
			this.minHeapUp(num)				// 否则扔到小根堆
		}
	} else {	// 元素比小根堆的堆顶元素小，需要扔到大根堆
		if len(this.maxArr) + 1 > len(this.minArr) {	// 如果将该元素扔到大根堆后，长度比小根堆要大，则不能直接扔进大根堆
			if num < this.maxArr[0] {					// 该元素比大根堆堆顶元素小
				this.minHeapUp(this.maxArr[0])			// 大根堆堆顶元素放到小根堆
				this.maxHeapDown(num)					// 把元素放到小根堆堆顶
			} else {
				this.minHeapUp(num)						// 该元素比大根堆堆顶大或相等，加入小根堆
			}
		} else {
			this.maxHeapUp(num)							// 如果将该元素扔到大根堆后，长度不比小根堆大，则直接扔进大根堆
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.maxArr) == len(this.minArr) {
		return (float64(this.maxArr[0]) + float64(this.minArr[0]) / 2)
	} else {
		return float64(this.minArr[0])
	}
}

//minHeapUp 小顶堆上浮
func (this *MedianFinder) minHeapUp(val int) {
	this.minArr = append(this.minArr, val)
	//
	i := len(this.minArr) - 1
	for i > 0 && this.minArr[i] < this.minArr[(i-1)/2] {
		this.minArr[i], this.minArr[(i-1)/2] = this.minArr[(i-1)/2], this.minArr[i]
		i = (i-1) / 2
	}
}

//maxHeapUp 大顶堆上浮
func (this *MedianFinder) maxHeapUp(val int) {
	this.maxArr = append(this.maxArr, val)
	i := len(this.maxArr) - 1
	for i > 0 && this.maxArr[i] < this.maxArr[(i-1)/2] {
		this.maxArr[i], this.maxArr[(i-1)/2] = this.maxArr[(i-1)/2], this.maxArr[i]
		i = (i-1) / 2
	}
}

//minHeapDown 小根堆下降
func (this *MedianFinder) minHeapDown(val int) {
	index := 0
	l := 2 * index + 1

	minest := 0
	this.minArr[0] = val

	for l < len(this.minArr) {
		if l + 1 < len(this.minArr) && this.minArr[l] > this.minArr[l+1] {
			minest = l + 1
		} else {
			minest = l
		}

		if this.minArr[index] < this.minArr[minest] {
			minest = index
		}

		if index == minest {
			break
		}

		this.minArr[index], this.minArr[minest] = this.minArr[minest], this.minArr[index]
		minest = index
		l = index * 2 + 1
	}
}

// 大根堆更新堆顶元素后的操作
func (this *MedianFinder) maxHeapDown(val int) {
    index := 0
    l := index * 2 + 1

    maxest := 0
    this.maxArr[0] = val

    for l < len(this.maxArr) {
        if l + 1 < len(this.maxArr) && this.maxArr[l] < this.maxArr[l+1] {
            maxest = l + 1
        } else {
            maxest = l
        }

        if this.maxArr[index] > this.maxArr[maxest] {
            maxest = index
        }

        if index == maxest {
            break
		}
		
		this.maxArr[index], this.maxArr[maxest] = this.maxArr[maxest], this.maxArr[index]
        index = maxest
        l = index * 2 + 1
    }
}